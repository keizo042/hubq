package di

import (
	"context"

	"github.com/keizo042/hubq/github"
	"github.com/morikuni/failure"
)

func (c *Container) InjectGithub(ctx context.Context) (github.Client, error) {
	httpClient, err := c.InjectHTTPClient(ctx)
	if err != nil {
		return nil, failure.Wrap(err)
	}
	client, err := github.New(httpClient)
	if err != nil {
		return nil, failure.Wrap(err)
	}
	return client, nil
}
