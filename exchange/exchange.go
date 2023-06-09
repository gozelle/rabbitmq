package exchange

import "github.com/gozelle/amqp"

func WithName(name string) Option {
	return func(c *config) {
		c.name = name
	}
}

func WithKind(kind string) Option {
	return func(c *config) {
		c.kind = kind
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

func WithInternal(v bool) Option {
	return func(c *config) {
		c.internal = v
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

type Option func(c *config)

type config struct {
	name       string
	kind       string
	durable    bool
	autoDelete bool
	internal   bool
	noWait     bool
	args       amqp.Table
}

func (c config) Name() string {
	return c.name
}

func (c config) Kind() string {
	return c.kind
}

func (c config) Durable() bool {
	return c.durable
}

func (c config) AutoDelete() bool {
	return c.autoDelete
}

func (c config) Internal() bool {
	return c.internal
}

func (c config) NoWait() bool {
	return c.noWait
}

func (c config) Args() amqp.Table {
	return c.args
}

func NewExchange(option ...Option) *Exchange {
	c := &config{}
	for _, v := range option {
		v(c)
	}
	return &Exchange{
		config: c,
	}
}

type Exchange struct {
	*config
}
