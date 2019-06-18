package di

import (
	"github.com/keizo042/hubq/config"
	"github.com/urfave/cli"
)

type (
	Container struct {
		Config *config.Config

		Cache struct {
			CLI *cli.App
		}
	}
)

func NewContainer(c *config.Config) (*Container, error) {
	return &Container{}, nil
}
