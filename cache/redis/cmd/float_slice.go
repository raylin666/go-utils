package cmd

import (
	"github.com/go-redis/redis/v8"
)

type FloatSliceCmd struct {
	CMD *redis.FloatSliceCmd
}

func NewFloatSliceCMD(cmd *redis.FloatSliceCmd) *FloatSliceCmd {
	return &FloatSliceCmd{cmd}
}

func (c *FloatSliceCmd) Result() ([]float64, error) {
	return c.CMD.Result()
}

func (c *FloatSliceCmd) String() string {
	return c.CMD.String()
}

func (c *FloatSliceCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *FloatSliceCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *FloatSliceCmd) Val() []float64 {
	return c.CMD.Val()
}

func (c *FloatSliceCmd) Err() error {
	return c.CMD.Err()
}

func (c *FloatSliceCmd) SetVal(val []float64) {
	c.CMD.SetVal(val)
}

func (c *FloatSliceCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}