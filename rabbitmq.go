package rabbitmq

import (
	"fmt"
	"github.com/gozelle/amqp"
	"github.com/gozelle/rabbitmq/consumer"
	"github.com/gozelle/rabbitmq/producer"
)

func NewRabbitMQ[T any](addr string) *RabbitMQ[T] {
	return &RabbitMQ[T]{addr: addr}
}

type RabbitMQ[T any] struct {
	addr string
}

func (r RabbitMQ[T]) openConn() (conn *amqp.Connection, err error) {
	conn, err = amqp.Dial(r.addr)
	if err != nil {
		return
	}
	return
}

func (r RabbitMQ[T]) openChannel(conn *amqp.Connection) (ch *amqp.Channel, err error) {
	ch, err = conn.Channel()
	if err != nil {
		err = fmt.Errorf("open channel error: %s", err)
		return
	}
	return
}

func (r RabbitMQ[T]) NewConsumer(options ...consumer.Option) (p *consumer.Consumer[T], err error) {
	c := &consumer.Config{}
	for _, v := range options {
		v(c)
	}
	conn, err := r.openConn()
	if err != nil {
		return
	}
	ch, err := r.openChannel(conn)
	if err != nil {
		return
	}
	p = consumer.NewConsumer[T](conn, ch, c)
	return
}

func (r RabbitMQ[T]) NewProducer(options ...producer.Option) (p *producer.Producer[T], err error) {

	c := &producer.Config{}
	for _, v := range options {
		v(c)
	}

	conn, err := r.openConn()
	if err != nil {
		return
	}

	ch, err := r.openChannel(conn)
	if err != nil {
		return
	}

	if c.Exchange() == nil {
		err = fmt.Errorf("expect exchange define use WithExchange")
		return
	}
	err = ch.ExchangeDeclare(
		c.Exchange().Name(),
		c.Exchange().Kind(),
		c.Exchange().Durable(),
		c.Exchange().AutoDelete(),
		c.Exchange().Internal(),
		c.Exchange().NoWait(),
		c.Exchange().Args(),
	)
	if err != nil {
		err = fmt.Errorf("exchange declare error: %s", err)
		return
	}

	for _, v := range c.Queues() {
		var vv amqp.Queue
		vv, err = ch.QueueDeclare(
			v.Name(),
			v.Durable(),
			v.AutoDelete(),
			v.Exclusive(),
			v.NoWait(),
			v.Args(),
		)
		if err != nil {
			err = fmt.Errorf("queue declare error: %s", err)
			return
		}

		err = ch.QueueBind(
			vv.Name,
			v.BindKey(),
			c.Exchange().Name(),
			v.NoWait(),
			v.Args(),
		)
		if err != nil {
			err = fmt.Errorf("queue bind error: %s", err)
			return
		}
	}

	p = producer.NewProducer[T](conn, ch, c.Exchange().Name())

	return
}
