package di

import (
	"github.com/urfave/cli"
)

func (c *Container) InjectCli() (*cli.App, error) {
	if c.Cache.CLI == nil {
		app := cli.NewApp()
		c.Cache.CLI = app
	}
	return c.Cache.CLI, nil
}
