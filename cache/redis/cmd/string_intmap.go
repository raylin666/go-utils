package cmd

import (
	"github.com/go-redis/redis/v8"
)

type StringIntMapCmd struct {
	CMD *redis.StringIntMapCmd
}

func NewStringIntMapCMD(cmd *redis.StringIntMapCmd) *StringIntMapCmd {
	return &StringIntMapCmd{cmd}
}

func (c *StringIntMapCmd) Result() (map[string]int64, error) {
	return c.CMD.Result()
}

func (c *StringIntMapCmd) String() string {
	return c.CMD.String()
}

func (c *StringIntMapCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *StringIntMapCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *StringIntMapCmd) Val() map[string]int64 {
	return c.CMD.Val()
}

func (c *StringIntMapCmd) Err() error {
	return c.CMD.Err()
}

func (c *StringIntMapCmd) SetVal(val map[string]int64) {
	c.CMD.SetVal(val)
}

func (c *StringIntMapCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}