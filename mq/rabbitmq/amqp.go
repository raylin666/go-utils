package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Client struct {
	Conn *amqp.Connection
}

type Channel struct {
	*amqp.Channel
}

type Parameters struct {
	// 队列名称
	QueueName  string
	// 队列 Key, 一般情况和队列名称一样即可
	QueueKey   string
	// 是否持久化, 即服务器重新启动后继续存在
	Durable    bool
	// 是否自动删除，关闭通道后就将删除
	AutoDelete bool
	// 是否独占队列
	Exclusive  bool
	// 是否等待, 无需等待服务器的响应
	NoWait     bool
	// 其他参数
	Args       amqp.Table

	// 交换机类型 fanout | direct | headers | topic | x-delayed-message
	ExchangeType string
	// 交换机名称
	ExchangeName string
	// 内部的, 不应向代理用户公开的交换间拓扑
	Internal     bool

	// 强制标志
	Mandatory    bool
	// 套接字
	Immediate    bool

	// 消费者标识, 该字符串是唯一的，适用于该频道上的所有消费者
	ConsumerId   string
	// 是否自动确认消息
	AutoAck		 bool
	// 暂不支持 noLocal 标志
	NoLocal		 bool
}

type Publishing struct {
	amqp.Publishing
}

type Options struct {
	Scheme   string
	User     string
	Password string
	Host     string
	Port     int
}

func New(opts *Options) (*Client, error) {
	opts = opts.init()
	url := fmt.Sprintf("%s://%s:%s@%s:%d",
		opts.Scheme,
		opts.User,
		opts.Password,
		opts.Host,
		opts.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	var client = new(Client)
	client.Conn = conn
	return client, nil
}

func (c *Client) Channel() (*Channel, error) {
	var ch = new(Channel)
	channel, err := c.Conn.Channel()
	if err != nil {
		return nil, err
	}

	ch.Channel = channel
	return ch, nil
}

func (ch *Channel) QueueDeclare(parameters *Parameters) (amqp.Queue, error) {
	return ch.Channel.QueueDeclare(
		parameters.QueueName,
		parameters.Durable,
		parameters.AutoDelete,
		parameters.Exclusive,
		parameters.NoWait,
		parameters.Args)
}

func (ch *Channel) QueueBind(parameters *Parameters) error {
	if parameters.QueueKey == "" {
		parameters.QueueKey = parameters.QueueName
	}

	return ch.Channel.QueueBind(
		parameters.QueueName,
		parameters.QueueKey,
		parameters.ExchangeName,
		parameters.NoWait,
		parameters.Args)
}

func (ch *Channel) ExchangeDeclare(parameters *Parameters) error {
	if parameters.ExchangeType == "" {
		parameters.ExchangeType = "direct"
	}

	return ch.Channel.ExchangeDeclare(
		parameters.ExchangeName,
		parameters.ExchangeType,
		parameters.Durable,
		parameters.AutoDelete,
		parameters.Internal,
		parameters.NoWait,
		parameters.Args)
}

func (ch *Channel) Publish(parameters *Parameters, publishing Publishing) error {
	if publishing.ContentType == "" {
		publishing.ContentType = "text/plain"
	}

	return ch.Channel.Publish(
		parameters.ExchangeName,
		parameters.QueueKey,
		parameters.Mandatory,
		parameters.Immediate,
		publishing.Publishing)
}

func (ch *Channel) Consume(parameters *Parameters) (<-chan amqp.Delivery, error) {
	return ch.Channel.Consume(
		parameters.QueueName,
		parameters.ConsumerId,
		parameters.AutoAck,
		parameters.Exclusive,
		parameters.NoLocal,
		parameters.NoWait,
		parameters.Args)
}

func (ch *Channel) Close() error {
	return ch.Channel.Close()
}

func (c *Client) PublishMessage(parameters *Parameters, publishing Publishing) (*amqp.Queue, error) {
	ch, err := c.Channel()
	if err != nil {
		return nil, err
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(parameters)
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(parameters)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(parameters)
	if err != nil {
		return nil, err
	}

	return &queue, ch.Publish(parameters, publishing)
}

func (c *Client) ConsumerMessage(parameters *Parameters) (<-chan amqp.Delivery, error) {
	ch, err := c.Channel()
	if err != nil {
		return nil, err
	}

	defer ch.Close()

	_, err = ch.QueueDeclare(parameters)
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(parameters)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(parameters)
	if err != nil {
		return nil, err
	}

	return ch.Consume(parameters)
}

func (c *Client) Close() error {
	return c.Conn.Close()
}

// init 初始化配置参数选项
func (opts *Options) init() *Options {
	if opts.Scheme == "" {
		opts.Scheme = "amqp"
	}

	if opts.User == "" {
		opts.User = "guest"
	}

	if opts.Password == "" {
		opts.Password = "guest"
	}

	if opts.Host == "" {
		opts.Host = "127.0.0.1"
	}

	if opts.Port == 0 {
		opts.Port = 5672
	}

	return opts
}
