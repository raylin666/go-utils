package cmd

import (
	"github.com/go-redis/redis/v8"
)

type ClusterSlotsCmd struct {
	CMD *redis.ClusterSlotsCmd
}

func NewClusterSlotsCMD(cmd *redis.ClusterSlotsCmd) *ClusterSlotsCmd {
	return &ClusterSlotsCmd{cmd}
}

func (c *ClusterSlotsCmd) Result() ([]redis.ClusterSlot, error) {
	return c.CMD.Result()
}

func (c *ClusterSlotsCmd) String() string {
	return c.CMD.String()
}

func (c *ClusterSlotsCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *ClusterSlotsCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *ClusterSlotsCmd) Val() []redis.ClusterSlot {
	return c.CMD.Val()
}

func (c *ClusterSlotsCmd) Err() error {
	return c.CMD.Err()
}

func (c *ClusterSlotsCmd) SetVal(val []redis.ClusterSlot) {
	c.CMD.SetVal(val)
}

func (c *ClusterSlotsCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}