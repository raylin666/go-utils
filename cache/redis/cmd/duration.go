package cmd

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type DurationCmd struct {
	CMD *redis.DurationCmd
}

func NewDurationCMD(cmd *redis.DurationCmd) *DurationCmd {
	return &DurationCmd{cmd}
}

func (c *DurationCmd) Result() (time.Duration, error) {
	return c.CMD.Result()
}

func (c *DurationCmd) String() string {
	return c.CMD.String()
}

func (c *DurationCmd) Name() string {
	return c.CMD.Name()
}

func (c *DurationCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *DurationCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *DurationCmd) Val() time.Duration {
	return c.CMD.Val()
}

func (c *DurationCmd) Err() error {
	return c.CMD.Err()
}

func (c *DurationCmd) SetVal(val time.Duration) {
	c.CMD.SetVal(val)
}

func (c *DurationCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}