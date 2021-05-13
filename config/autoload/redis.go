package autoload

import "time"

type Redis struct {
	Addr         string        `yaml:"Addr"`
	Port         int           `yaml:"Port"`
	Password     string        `yaml:"Password"`
	Db           int           `yaml:"Db"`
	MaxRetries   int           `yaml:"MaxRetries"`
	PoolSize     int           `yaml:"PoolSize"`
	PoolTimeout  time.Duration `yaml:"PoolTimeout"`
	MinIdleConns int           `yaml:"MinIdleConns"`
	IdleTimeout  time.Duration `yaml:"IdleTimeout"`
	DialTimeout  time.Duration `yaml:"DialTimeout"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}
