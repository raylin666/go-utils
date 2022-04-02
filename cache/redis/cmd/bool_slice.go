package cmd

import (
	"github.com/go-redis/redis/v8"
)

type BoolSliceCmd struct {
	CMD *redis.BoolSliceCmd
}

func NewBoolSliceCMD(cmd *redis.BoolSliceCmd) *BoolSliceCmd {
	return &BoolSliceCmd{cmd}
}

func (c *BoolSliceCmd) Result() ([]bool, error) {
	return c.CMD.Result()
}

func (c *BoolSliceCmd) String() string {
	return c.CMD.String()
}

func (c *BoolSliceCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *BoolSliceCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *BoolSliceCmd) Val() []bool {
	return c.CMD.Val()
}

func (c *BoolSliceCmd) Err() error {
	return c.CMD.Err()
}

func (c *BoolSliceCmd) SetVal(val []bool) {
	c.CMD.SetVal(val)
}

func (c *BoolSliceCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}