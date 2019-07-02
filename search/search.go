package search

import (
	"context"
	"fmt"
	"github.com/keizo042/hubq/errors"
	"github.com/keizo042/hubq/github"
	"github.com/morikuni/failure"
	"strings"
)

type (
	// TODO(keizo042): rename
	Search interface {
		Search(ctx context.Context, req *Request, options *Option) (*Response, error)
	}

	Request struct {
		RawQuery string
		Keyword  string
	}

	Response struct {
		Repositories []Repository
	}

	Repository struct {
		Name     string
		Author   string
		CloneURL string
	}

	Option struct {
	}
)

type (
	search struct {
		githubClient github.Client
	}
)

func New(githubClient github.Client) (Search, error) {
	return &search{
		githubClient: githubClient,
	}, nil
}

func (s *search) Search(ctx context.Context, req *Request, options *Option) (*Response, error) {
	if err := req.validate(); err != nil {
		return nil, failure.Wrap(err)
	}
	q, err := req.buildQuery()
	if err != nil {
		return nil, failure.Wrap(err)
	}
	res, err := s.githubClient.Search(ctx, q)
	if err != nil {
		return nil, failure.Wrap(err)
	}
	var repos []Repository
	for _, r := range res.Repositories {
		if r.SSHURL == nil {
			continue
		}
		repos = append(repos, Repository{
			Name:     r.GetName(),
			Author:   r.GetFullName(),
			CloneURL: *r.SSHURL,
		})

	}

	return &Response{
		Repositories: repos,
	}, nil
}

func (req *Request) validate() error {
	if req.RawQuery == "" &&
		req.Keyword == "" {
		e := fmt.Errorf("require some query parameter")
		return failure.Translate(e, errors.ConditionRequired)
	}
	return nil
}

func (req *Request) buildQuery() (string, error) {
	var b strings.Builder
	if _, err := b.WriteString(req.RawQuery); err != nil {
		return "", failure.Wrap(err)
	}
	if _, err := b.WriteString(" "); err != nil {
		return "", failure.Wrap(err)
	}
	if _, err := b.WriteString(req.Keyword); err != nil {
		return "", failure.Wrap(err)
	}
	if _, err := b.WriteString(" "); err != nil {
		return "", failure.Wrap(err)
	}
	return b.String(), nil
}
