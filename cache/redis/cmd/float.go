package cmd

import "github.com/go-redis/redis/v8"

type FloatCmd struct {
	CMD *redis.FloatCmd
}

func NewFloatCMD(cmd *redis.FloatCmd) *FloatCmd {
	return &FloatCmd{cmd}
}

func (c *FloatCmd) Result() (float64, error) {
	return c.CMD.Result()
}

func (c *FloatCmd) String() string {
	return c.CMD.String()
}

func (c *FloatCmd) Name() string {
	return c.CMD.Name()
}

func (c *FloatCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *FloatCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *FloatCmd) Val() float64 {
	return c.CMD.Val()
}

func (c *FloatCmd) Err() error {
	return c.CMD.Err()
}

func (c *FloatCmd) SetVal(val float64) {
	c.CMD.SetVal(val)
}

func (c *FloatCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}