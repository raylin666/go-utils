package cmd

import "github.com/go-redis/redis/v8"

type SlowLogCmd struct {
	CMD *redis.SlowLogCmd
}

func NewSlowLogCMD(cmd *redis.SlowLogCmd) *SlowLogCmd {
	return &SlowLogCmd{cmd}
}

func (c *SlowLogCmd) Result() ([]redis.SlowLog, error) {
	return c.CMD.Result()
}

func (c *SlowLogCmd) String() string {
	return c.CMD.String()
}

func (c *SlowLogCmd) Name() string {
	return c.CMD.Name()
}

func (c *SlowLogCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *SlowLogCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *SlowLogCmd) Val() []redis.SlowLog {
	return c.CMD.Val()
}

func (c *SlowLogCmd) Err() error {
	return c.CMD.Err()
}

func (c *SlowLogCmd) SetVal(val []redis.SlowLog) {
	c.CMD.SetVal(val)
}

func (c *SlowLogCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}