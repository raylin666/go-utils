package cmd

import "github.com/go-redis/redis/v8"

type GeoSearchLocationCmd struct {
	CMD *redis.GeoSearchLocationCmd
}

func NewGeoSearchLocationCMD(cmd *redis.GeoSearchLocationCmd) *GeoSearchLocationCmd {
	return &GeoSearchLocationCmd{cmd}
}

func (c *GeoSearchLocationCmd) Result() ([]redis.GeoLocation, error) {
	return c.CMD.Result()
}

func (c *GeoSearchLocationCmd) String() string {
	return c.CMD.String()
}

func (c *GeoSearchLocationCmd) Name() string {
	return c.CMD.Name()
}

func (c *GeoSearchLocationCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *GeoSearchLocationCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *GeoSearchLocationCmd) Val() []redis.GeoLocation {
	return c.CMD.Val()
}

func (c *GeoSearchLocationCmd) Err() error {
	return c.CMD.Err()
}

func (c *GeoSearchLocationCmd) SetVal(locations []redis.GeoLocation) {
	c.CMD.SetVal(locations)
}

func (c *GeoSearchLocationCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}