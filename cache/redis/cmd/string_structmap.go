package cmd

import (
	"github.com/go-redis/redis/v8"
)

type StringStructMapCmd struct {
	CMD *redis.StringStructMapCmd
}

func NewStringStructMapCMD(cmd *redis.StringStructMapCmd) *StringStructMapCmd {
	return &StringStructMapCmd{cmd}
}

func (c *StringStructMapCmd) Result() (map[string]struct{}, error) {
	return c.CMD.Result()
}

func (c *StringStructMapCmd) String() string {
	return c.CMD.String()
}

func (c *StringStructMapCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *StringStructMapCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *StringStructMapCmd) Val() map[string]struct{} {
	return c.CMD.Val()
}

func (c *StringStructMapCmd) Err() error {
	return c.CMD.Err()
}

func (c *StringStructMapCmd) SetVal(val map[string]struct{}) {
	c.CMD.SetVal(val)
}

func (c *StringStructMapCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}