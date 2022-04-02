package cmd

import "github.com/go-redis/redis/v8"

type GeoPosCmd struct {
	CMD *redis.GeoPosCmd
}

func NewGeoPosCMD(cmd *redis.GeoPosCmd) *GeoPosCmd {
	return &GeoPosCmd{cmd}
}

func (c *GeoPosCmd) Result() ([]*redis.GeoPos, error) {
	return c.CMD.Result()
}

func (c *GeoPosCmd) String() string {
	return c.CMD.String()
}

func (c *GeoPosCmd) Name() string {
	return c.CMD.Name()
}

func (c *GeoPosCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *GeoPosCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *GeoPosCmd) Val() []*redis.GeoPos {
	return c.CMD.Val()
}

func (c *GeoPosCmd) Err() error {
	return c.CMD.Err()
}

func (c *GeoPosCmd) SetVal(val []*redis.GeoPos) {
	c.CMD.SetVal(val)
}

func (c *GeoPosCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}