package cmd

import "github.com/go-redis/redis/v8"

type IntCmd struct {
	CMD *redis.IntCmd
}

func NewIntCMD(cmd *redis.IntCmd) *IntCmd {
	return &IntCmd{cmd}
}

func (c *IntCmd) Result() (int64, error) {
	return c.CMD.Result()
}

func (c *IntCmd) String() string {
	return c.CMD.String()
}

func (c *IntCmd) Uint64() (uint64, error) {
	return c.CMD.Uint64()
}

func (c *IntCmd) Name() string {
	return c.CMD.Name()
}

func (c *IntCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *IntCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *IntCmd) Val() int64 {
	return c.CMD.Val()
}

func (c *IntCmd) Err() error {
	return c.CMD.Err()
}

func (c *IntCmd) SetVal(val int64) {
	c.CMD.SetVal(val)
}

func (c *IntCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}