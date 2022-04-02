package cmd

import (
	"github.com/go-redis/redis/v8"
)

type StringStringMapCmd struct {
	CMD *redis.StringStringMapCmd
}

func NewStringStringMapCMD(cmd *redis.StringStringMapCmd) *StringStringMapCmd {
	return &StringStringMapCmd{cmd}
}

func (c *StringStringMapCmd) Result() (map[string]string, error) {
	return c.CMD.Result()
}

func (c *StringStringMapCmd) String() string {
	return c.CMD.String()
}

func (c *StringStringMapCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *StringStringMapCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *StringStringMapCmd) Val() map[string]string {
	return c.CMD.Val()
}

func (c *StringStringMapCmd) Err() error {
	return c.CMD.Err()
}

func (c *StringStringMapCmd) Scan(dest interface{}) error {
	return c.CMD.Scan(dest)
}

func (c *StringStringMapCmd) SetVal(val map[string]string) {
	c.CMD.SetVal(val)
}

func (c *StringStringMapCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}