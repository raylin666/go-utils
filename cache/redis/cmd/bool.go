package cmd

import "github.com/go-redis/redis/v8"

type BoolCmd struct {
	CMD *redis.BoolCmd
}

func NewBoolCMD(cmd *redis.BoolCmd) *BoolCmd {
	return &BoolCmd{cmd}
}

func (c *BoolCmd) Result() (bool, error) {
	return c.CMD.Result()
}

func (c *BoolCmd) String() string {
	return c.CMD.String()
}

func (c *BoolCmd) Name() string {
	return c.CMD.Name()
}

func (c *BoolCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *BoolCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *BoolCmd) Val() bool {
	return c.CMD.Val()
}

func (c *BoolCmd) Err() error {
	return c.CMD.Err()
}

func (c *BoolCmd) SetVal(val bool) {
	c.CMD.SetVal(val)
}

func (c *BoolCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}