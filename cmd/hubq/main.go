package main

import (
	"fmt"
	"os"
	"path/filepath"

	hubcli "github.com/keizo042/hubq/cli"
	"github.com/urfave/cli"
)

const (
	exitSuccess = iota
	exitFailure
)

var (
	DefaultConfigPath = "/.hubq/config"
)

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(argv []string) int {
	c := cli.NewApp()
	path := filepath.Join(os.Getenv("HOME"), DefaultConfigPath)
	commands, err := hubcli.New(path)
	if err != nil {
		return exitFailure
	}
	c.Commands = []cli.Command{
		commands.SearchCommand(),
	}

	if err := c.Run(argv); err != nil {
		fmt.Println(err)
		return exitFailure
	}
	return exitSuccess
}
