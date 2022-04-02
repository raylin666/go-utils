package cmd

import (
	"github.com/go-redis/redis/v8"
)

type ZSliceCmd struct {
	CMD *redis.ZSliceCmd
}

func NewZSliceCMD(cmd *redis.ZSliceCmd) *ZSliceCmd {
	return &ZSliceCmd{cmd}
}

func (c *ZSliceCmd) Result() ([]redis.Z, error) {
	return c.CMD.Result()
}

func (c *ZSliceCmd) String() string {
	return c.CMD.String()
}

func (c *ZSliceCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *ZSliceCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *ZSliceCmd) Val() []redis.Z {
	return c.CMD.Val()
}

func (c *ZSliceCmd) Err() error {
	return c.CMD.Err()
}

func (c *ZSliceCmd) SetVal(val []redis.Z) {
	c.CMD.SetVal(val)
}

func (c *ZSliceCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}