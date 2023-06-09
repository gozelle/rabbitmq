package queue

import "github.com/gozelle/amqp"

func WithName(name string) Option {
	return func(c *config) {
		c.name = name
	}
}

func WithDurable(v bool) Option {
	return func(c *config) {
		c.durable = v
	}
}

func WithAutoDelete(v bool) Option {
	return func(c *config) {
		c.autoDelete = v
	}
}

func WithExclusive(v bool) Option {
	return func(c *config) {
		c.exclusive = v
	}
}

func WithNoWait(v bool) Option {
	return func(c *config) {
		c.noWait = v
	}
}

func WithArgs(v amqp.Table) Option {
	return func(c *config) {
		c.args = v
	}
}

func WithKey(key string) Option {
	return func(c *config) {
		c.bindKey = key
	}
}

type Option func(c *config)

type config struct {
	name       string
	bindKey    string
	durable    bool
	autoDelete bool
	exclusive  bool
	noWait     bool
	args       amqp.Table
}

func (c config) BindKey() string {
	return c.bindKey
}

func (c config) Name() string {
	return c.name
}

func (c config) Durable() bool {
	return c.durable
}

func (c config) AutoDelete() bool {
	return c.autoDelete
}

func (c config) Exclusive() bool {
	return c.exclusive
}

func (c config) NoWait() bool {
	return c.noWait
}

func (c config) Args() amqp.Table {
	return c.args
}

func NewQueue(option ...Option) *Queue {
	c := &config{}
	for _, v := range option {
		v(c)
	}
	return &Queue{
		config: c,
	}
}

type Queue struct {
	*config
}
