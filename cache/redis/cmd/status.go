package cmd

import "github.com/go-redis/redis/v8"

type StatusCmd struct {
	CMD *redis.StatusCmd
}

func NewStatusCMD(cmd *redis.StatusCmd) *StatusCmd {
	return &StatusCmd{cmd}
}

func (c *StatusCmd) Result() (string, error) {
	return c.CMD.Result()
}

func (c *StatusCmd) String() string {
	return c.CMD.String()
}

func (c *StatusCmd) Name() string {
	return c.CMD.Name()
}

func (c *StatusCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *StatusCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *StatusCmd) Val() string {
	return c.CMD.Val()
}

func (c *StatusCmd) Err() error {
	return c.CMD.Err()
}

func (c *StatusCmd) SetVal(val string) {
	c.CMD.SetVal(val)
}

func (c *StatusCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}