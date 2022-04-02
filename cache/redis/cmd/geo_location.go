package cmd

import "github.com/go-redis/redis/v8"

type GeoLocationCmd struct {
	CMD *redis.GeoLocationCmd
}

func NewGeoLocationCMD(cmd *redis.GeoLocationCmd) *GeoLocationCmd {
	return &GeoLocationCmd{cmd}
}

func (c *GeoLocationCmd) Result() ([]redis.GeoLocation, error) {
	return c.CMD.Result()
}

func (c *GeoLocationCmd) String() string {
	return c.CMD.String()
}

func (c *GeoLocationCmd) Name() string {
	return c.CMD.Name()
}

func (c *GeoLocationCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *GeoLocationCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *GeoLocationCmd) Val() []redis.GeoLocation {
	return c.CMD.Val()
}

func (c *GeoLocationCmd) Err() error {
	return c.CMD.Err()
}

func (c *GeoLocationCmd) SetVal(locations []redis.GeoLocation) {
	c.CMD.SetVal(locations)
}

func (c *GeoLocationCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}