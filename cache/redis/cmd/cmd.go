package cmd

import "github.com/go-redis/redis/v8"

type Cmd struct {
	CMD *redis.Cmd
}

func NewCMD(cmd *redis.Cmd) *Cmd {
	return &Cmd{cmd}
}

func (c *Cmd) Result() (interface{}, error) {
	return c.CMD.Result()
}

func (c *Cmd) String() string {
	return c.CMD.String()
}

func (c *Cmd) StringSlice() ([]string, error) {
	return c.CMD.StringSlice()
}

func (c *Cmd) Uint64() (uint64, error) {
	return c.CMD.Uint64()
}

func (c *Cmd) Uint64Slice() ([]uint64, error) {
	return c.CMD.Uint64Slice()
}

func (c *Cmd) Int64() (int64, error) {
	return c.CMD.Int64()
}

func (c *Cmd) Int64Slice() ([]int64, error) {
	return c.CMD.Int64Slice()
}

func (c *Cmd) Int() (int, error) {
	return c.CMD.Int()
}

func (c *Cmd) Float64() (float64, error) {
	return c.CMD.Float64()
}

func (c *Cmd) Float64Slice() ([]float64, error) {
	return c.CMD.Float64Slice()
}

func (c *Cmd) Float32() (float32, error) {
	return c.CMD.Float32()
}

func (c *Cmd) Float32Slice() ([]float32, error) {
	return c.CMD.Float32Slice()
}

func (c *Cmd) Bool() (bool, error) {
	return c.CMD.Bool()
}

func (c *Cmd) BoolSlice() ([]bool, error) {
	return c.CMD.BoolSlice()
}

func (c *Cmd) Slice() ([]interface{}, error) {
	return c.CMD.Slice()
}

func (c *Cmd) Text() (string, error) {
	return c.CMD.Text()
}

func (c *Cmd) Name() string {
	return c.CMD.Name()
}

func (c *Cmd) FullName() string {
	return c.CMD.FullName()
}

func (c *Cmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *Cmd) Val() interface{} {
	return c.CMD.Val()
}

func (c *Cmd) Err() error {
	return c.CMD.Err()
}

func (c *Cmd) SetVal(val bool) {
	c.CMD.SetVal(val)
}

func (c *Cmd) SetErr(e error) {
	c.CMD.SetErr(e)
}