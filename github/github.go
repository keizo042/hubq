package github

import (
	"context"
	github "github.com/google/go-github/github"
	"github.com/morikuni/failure"
	"net/http"
)

type (
	Client interface {
		Search(context.Context, string) (*github.RepositoriesSearchResult, error)
	}

	client struct {
		githubClient *github.Client
	}
)

func New(httpClient *http.Client) (Client, error) {
	githubClient := github.NewClient(httpClient)
	return &client{
		githubClient: githubClient,
	}, nil
}

func (c *client) Search(ctx context.Context, query string) (*github.RepositoriesSearchResult, error) {
	result, _, err := c.githubClient.Search.Repositories(ctx, query, nil)
	if err != nil {
		return nil, failure.Wrap(err)
	}
	return result, nil
}
