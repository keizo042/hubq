package main

import (
	"fmt"
	"os"
	"path/filepath"

	hubcli "github.com/keizo042/hubq/cli"
	"github.com/morikuni/failure"
	"github.com/urfave/cli"
)

const (
	exitSuccess = iota
	exitFailure
)

var (
	Version           string
	DefaultConfigPath = "/.hubq/config"
)

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(argv []string) int {
	c := cli.NewApp()
	c.Version = Version
	path := filepath.Join(os.Getenv("HOME"), DefaultConfigPath)
	commands, err := hubcli.New(path)
	if err != nil {
		return exitFailure
	}
	c.Commands = []cli.Command{
		commands.SearchCommand(),
	}

	if err := c.Run(argv); err != nil {
		code, ok := failure.CodeOf(err)
		if ok {
			fmt.Println(code)
			return exitFailure
		}
		fmt.Printf("%+v\n", err)
		return exitFailure
	}
	return exitSuccess
}
