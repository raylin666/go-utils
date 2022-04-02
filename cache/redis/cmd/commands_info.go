package cmd

import "github.com/go-redis/redis/v8"

type CommandsInfoCmd struct {
	CMD *redis.CommandsInfoCmd
}

func NewCommandsInfoCMD(cmd *redis.CommandsInfoCmd) *CommandsInfoCmd {
	return &CommandsInfoCmd{cmd}
}

func (c *CommandsInfoCmd) Result() (map[string]*redis.CommandInfo, error) {
	return c.CMD.Result()
}

func (c *CommandsInfoCmd) String() string {
	return c.CMD.String()
}

func (c *CommandsInfoCmd) Name() string {
	return c.CMD.Name()
}

func (c *CommandsInfoCmd) FullName() string {
	return c.CMD.FullName()
}

func (c *CommandsInfoCmd) Args() []interface{} {
	return c.CMD.Args()
}

func (c *CommandsInfoCmd) Val() map[string]*redis.CommandInfo {
	return c.CMD.Val()
}

func (c *CommandsInfoCmd) Err() error {
	return c.CMD.Err()
}

func (c *CommandsInfoCmd) SetVal(val map[string]*redis.CommandInfo) {
	c.CMD.SetVal(val)
}

func (c *CommandsInfoCmd) SetErr(e error) {
	c.CMD.SetErr(e)
}
