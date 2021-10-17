package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

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
	ConsumerTag  string
	// 是否自动确认消息
	AutoAck		 bool
	// 暂不支持 noLocal 标志
	NoLocal		 bool
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
