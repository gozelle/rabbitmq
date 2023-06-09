package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/gozelle/amqp"
)

func WithQueue(name string) Option {
	return func(c *Config) {
		c.queue = name
	}
}

func WithConsumerName(name string) Option {
	return func(c *Config) {
		c.consumerName = name
	}
}

func WithAutoAck(v bool) Option {
	return func(c *Config) {
		c.autoAck = v
	}
}

func WithExclusive(v bool) Option {
	return func(c *Config) {
		c.exclusive = v
	}
}

func WithNoLocal(v bool) Option {
	return func(c *Config) {
		c.noLocal = v
	}
}

func WithNoWait(v bool) Option {
	return func(c *Config) {
		c.noWait = v
	}
}

func WithArgs(v amqp.Table) Option {
	return func(c *Config) {
		c.args = v
	}
}

type Option func(c *Config)

type Config struct {
	queue        string
	consumerName string
	autoAck      bool
	exclusive    bool
	noLocal      bool
	noWait       bool
	args         amqp.Table
}

func NewConsumer[T any](conn *amqp.Connection, ch *amqp.Channel, config *Config) *Consumer[T] {
	return &Consumer[T]{
		conn:   conn,
		ch:     ch,
		config: config,
	}
}

type Consumer[T any] struct {
	conn   *amqp.Connection
	ch     *amqp.Channel
	config *Config
}

func (p Consumer[T]) Close() {
	_ = p.ch.Close()
	_ = p.conn.Close()
}

func (p Consumer[T]) Consume(handler func(message T) error) (err error) {
	c := p.config
	messages, err := p.ch.Consume(
		c.queue,
		c.consumerName,
		c.autoAck,
		c.exclusive,
		c.noLocal,
		c.noWait,
		c.args,
	)
	if err != nil {
		err = fmt.Errorf("queue consume error: %s", err)
		return
	}
	
	for d := range messages {
		m := new(T)
		err = json.Unmarshal(d.Body, m)
		if err != nil {
			err = fmt.Errorf("unmarshal data error: %s, data: %s", err, string(d.Body))
			return
		}
		err = handler(*m)
		if err != nil {
			err = fmt.Errorf("handler message error: %s", err)
			return
		}
		if !c.autoAck {
			err = d.Ack(false)
			if err != nil {
				err = fmt.Errorf("ack error: %s", err)
				return
			}
		}
	}
	
	return
}
