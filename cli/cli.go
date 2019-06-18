package cli

import (
	"github.com/morikuni/failure"
	"github.com/urfave/cli"
)

type Commands struct {
	ConfigPath string
}

func New(configPath string) (*Commands, error) {
	return &Commands{
		ConfigPath: configPath,
	}, nil
}

func (com *Commands) Action(c *cli.Context) error {
	return failure.Wrap(nil)
}
