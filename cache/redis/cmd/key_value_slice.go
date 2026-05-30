package cmd

import (
	"github.com/redis/go-redis/v9"
)

type KeyValueSliceCmd struct {
	CMD *redis.KeyValueSliceCmd
}

func NewKeyValueSliceCMD(cmd *redis.KeyValueSliceCmd) *KeyValueSliceCmd {
	return &KeyValueSliceCmd{cmd}
}

func (c *KeyValueSliceCmd) Result() ([]redis.KeyValue, error) {
	return c.CMD.Result()
}

func (c *KeyValueSliceCmd) String() string {
	return c.CMD.String()
}

func (c *KeyValueSliceCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *KeyValueSliceCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *KeyValueSliceCmd) Val() []redis.KeyValue {
	return c.CMD.Val()
}

func (c *KeyValueSliceCmd) Err() error {
	return c.CMD.Err()
}

func (c *KeyValueSliceCmd) SetVal(val []redis.KeyValue) {
	c.CMD.SetVal(val)
}

func (c *KeyValueSliceCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}