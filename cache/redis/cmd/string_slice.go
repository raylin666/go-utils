package cmd

import (
	"github.com/go-redis/redis/v8"
)

type StringSliceCmd struct {
	CMD *redis.StringSliceCmd
}

func NewStringSliceCMD(cmd *redis.StringSliceCmd) *StringSliceCmd {
	return &StringSliceCmd{cmd}
}

func (c *StringSliceCmd) Result() ([]string, error) {
	return c.CMD.Result()
}

func (c *StringSliceCmd) String() string {
	return c.CMD.String()
}

func (c *StringSliceCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *StringSliceCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *StringSliceCmd) Val() []string {
	return c.CMD.Val()
}

func (c *StringSliceCmd) Err() error {
	return c.CMD.Err()
}

func (c *StringSliceCmd) ScanSlice(container interface{}) error {
	return c.CMD.ScanSlice(container)
}

func (c *StringSliceCmd) SetVal(val []string) {
	c.CMD.SetVal(val)
}

func (c *StringSliceCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}