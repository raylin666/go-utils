package cmd

import "github.com/go-redis/redis/v8"

type ScanCmd struct {
	CMD *redis.ScanCmd
}

func NewScanCMD(cmd *redis.ScanCmd) *ScanCmd {
	return &ScanCmd{cmd}
}

func (c *ScanCmd) Result() (keys []string, cursor uint64, err error) {
	return c.CMD.Result()
}

func (c *ScanCmd) String() string {
	return c.CMD.String()
}

func (c *ScanCmd) Name() string {
	return c.CMD.Name()
}

func (c *ScanCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *ScanCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *ScanCmd) Iterator() *redis.ScanIterator {
	return c.CMD.Iterator()
}

func (c *ScanCmd) Val() (keys []string, cursor uint64) {
	return c.CMD.Val()
}

func (c *ScanCmd) Err() error {
	return c.CMD.Err()
}

func (c *ScanCmd) SetVal(page []string, cursor uint64) {
	c.CMD.SetVal(page, cursor)
}

func (c *ScanCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}