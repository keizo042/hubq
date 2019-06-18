package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/morikuni/failure"
)

type Config struct {
	GithubToken string `toml:"github_token"`
}

func Parse(path string) (*Config, error) {
	var c Config
	if _, err := toml.DecodeFile(path, &c); err != nil {
		return nil, failure.Wrap(err)
	}
	return &c, nil
}

func (c *Config) validate() error {
	if c.GithubToken == "" {
		return fmt.Errorf("github_token is required")
	}
	return nil
}
