package cmd

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type TimeCmd struct {
	CMD *redis.TimeCmd
}

func NewTimeCMD(cmd *redis.TimeCmd) *TimeCmd {
	return &TimeCmd{cmd}
}

func (c *TimeCmd) Result() (time.Time, error) {
	return c.CMD.Result()
}

func (c *TimeCmd) String() string {
	return c.CMD.String()
}

func (c *TimeCmd) Name() string {
	return c.CMD.Name()
}

func (c *TimeCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *TimeCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *TimeCmd) Val() time.Time {
	return c.CMD.Val()
}

func (c *TimeCmd) Err() error {
	return c.CMD.Err()
}

func (c *TimeCmd) SetVal(val time.Time) {
	c.CMD.SetVal(val)
}

func (c *TimeCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}