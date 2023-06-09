# RabbitMQ

对 RabbitMQ 操作的封装，简化接入的工作量

发布消息:

```go
r := rabbitmq.NewRabbitMQ[string]("amqp://root:123456@127.0.01:5672")
	p, err := r.NewProducer(
		producer.WithExchange(
			exchange.NewExchange(
				exchange.WithName("test_exchange"),     // 指定交换器
				exchange.WithKind(amqp.ExchangeFanout), // 交换器类型
				exchange.WithDurable(true),             // 交换器持久化
			),
		),
		producer.WithQueues(
			queue.NewQueue(
				queue.WithName("test_exchange_q1"), // 队列名称
				queue.WithDurable(true),            // 队列持久化
			),
		),
	)
	require.NoError(t, err)
	
	err = p.Publish(message.NewMessage(Message{Id: 123}).WitPersistent(true))
	require.NoError(t, err)
```

消费消息：

```go
	r := rabbitmq.NewRabbitMQ[*Message]("amqp://root:123456@127.0.01:5672")
	c, err := r.NewConsumer(
		consumer.WithQueue("test_exchange_q1"), // 指定消费的队列
	)
	require.NoError(t, err)
	
	err = c.Consume(func(message *Message) error {
		fmt.Println(message)
		return nil
	})
	require.NoError(t, err)
```