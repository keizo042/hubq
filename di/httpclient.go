package di

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

func (c *Container) InjectHTTPClient(ctx context.Context) (*http.Client, error) {
	return oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: c.Config.GithubToken,
	})), nil
}
