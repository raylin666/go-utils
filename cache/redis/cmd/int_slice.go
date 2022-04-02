package cmd

import (
	"github.com/go-redis/redis/v8"
)

type IntSliceCmd struct {
	CMD *redis.IntSliceCmd
}

func NewIntSliceCMD(cmd *redis.IntSliceCmd) *IntSliceCmd {
	return &IntSliceCmd{cmd}
}

func (c *IntSliceCmd) Result() ([]int64, error) {
	return c.CMD.Result()
}

func (c *IntSliceCmd) String() string {
	return c.CMD.String()
}

func (c *IntSliceCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *IntSliceCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *IntSliceCmd) Val() []int64 {
	return c.CMD.Val()
}

func (c *IntSliceCmd) Err() error {
	return c.CMD.Err()
}

func (c *IntSliceCmd) SetVal(val []int64) {
	c.CMD.SetVal(val)
}

func (c *IntSliceCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}