package cmd

import (
	"github.com/go-redis/redis/v8"
)

type SliceCmd struct {
	CMD *redis.SliceCmd
}

func NewSliceCMD(cmd *redis.SliceCmd) *SliceCmd {
	return &SliceCmd{cmd}
}

func (c *SliceCmd) Result() ([]interface{}, error) {
	return c.CMD.Result()
}

func (c *SliceCmd) String() string {
	return c.CMD.String()
}

func (c *SliceCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *SliceCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *SliceCmd) Val() []interface{} {
	return c.CMD.Val()
}

func (c *SliceCmd) Err() error {
	return c.CMD.Err()
}

func (c *SliceCmd) Scan(dst interface{}) error {
	return c.CMD.Scan(dst)
}

func (c *SliceCmd) SetVal(val []interface{}) {
	c.CMD.SetVal(val)
}

func (c *SliceCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}