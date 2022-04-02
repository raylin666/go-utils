package cmd

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type StringCmd struct {
	CMD *redis.StringCmd
}

func NewStringCMD(cmd *redis.StringCmd) *StringCmd {
	return &StringCmd{cmd}
}

func (c *StringCmd) Result() (string, error) {
	return c.CMD.Result()
}

func (c *StringCmd) String() string {
	return c.CMD.String()
}

func (c *StringCmd) Float64() (float64, error) {
	return c.CMD.Float64()
}

func (c *StringCmd) Float32() (float32, error) {
	return c.CMD.Float32()
}

func (c *StringCmd) Bool() (bool, error) {
	return c.CMD.Bool()
}

func (c *StringCmd) Int64() (int64, error) {
	return c.CMD.Int64()
}

func (c *StringCmd) Uint64() (uint64, error) {
	return c.CMD.Uint64()
}

func (c *StringCmd) Int() (int, error) {
	return c.CMD.Int()
}

func (c *StringCmd) Bytes() ([]byte, error) {
	return c.CMD.Bytes()
}

func (c *StringCmd) Time() (time.Time, error) {
	return c.CMD.Time()
}

func (c *StringCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *StringCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *StringCmd) Val() string {
	return c.CMD.Val()
}

func (c *StringCmd) Err() error {
	return c.CMD.Err()
}

func (c *StringCmd) Scan(val interface{}) error {
	return c.CMD.Scan(val)
}

func (c *StringCmd) SetVal(val string) {
	c.CMD.SetVal(val)
}

func (c *StringCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}