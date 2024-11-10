package producer

import (
	"encoding/json"
	"fmt"
	"github.com/gozelle/amqp"
	"github.com/gozelle/rabbitmq/exchange"
	"github.com/gozelle/rabbitmq/message"
	"github.com/gozelle/rabbitmq/queue"
)

func WithExchange(e *exchange.Exchange) Option {
	return func(c *Config) {
		c.exchange = e
	}
}

func WithQueues(queues ...*queue.Queue) Option {
	return func(c *Config) {
		c.queues = append(c.queues, queues...)
	}
}

type Option func(c *Config)

type Config struct {
	exchange *exchange.Exchange
	queues   []*queue.Queue
}

func (c Config) Exchange() *exchange.Exchange {
	return c.exchange
}

func (c Config) Queues() []*queue.Queue {
	return c.queues
}

func NewProducer(conn *amqp.Connection, ch *amqp.Channel, exchange string) *Producer {
	return &Producer{conn: conn, ch: ch, exchange: exchange}
}

type Producer struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	exchange string
}

func (p Producer) Close() {
	_ = p.ch.Close()
	_ = p.conn.Close()
}

func (p Producer) Publish(m *message.Message) (err error) {

	data, err := json.Marshal(m.Value())
	if err != nil {
		err = fmt.Errorf("marshal message value error: %s", err)
		return
	}

	contentType := "text/plain"
	if m.ContentType() != "" {
		contentType = m.ContentType()
	}
	deliveryMode := amqp.Transient
	if m.Persistent() {
		deliveryMode = amqp.Persistent
	}

	ex := p.exchange
	if m.Exchange() != nil {
		ex = *m.Exchange()
	}

	err = p.ch.Publish(
		ex,
		m.RouteKey(),
		m.Mandatory(),
		m.Immediate(),
		amqp.Publishing{
			ContentType:  contentType,
			Body:         data,
			DeliveryMode: deliveryMode,
		},
	)
	if err != nil {
		err = fmt.Errorf("publish message error: %s", err)
		return
	}

	return
}
