package search

import (
	"context"
	"fmt"
	"github.com/keizo042/hubq/github"
	"github.com/morikuni/failure"
)

type (
	// TODO(keizo042): rename
	Search interface {
		Search(ctx context.Context, req *Request, options *Option) (*Response, error)
	}

	Request struct {
		Keyword string
	}

	Response struct {
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
	if _, err := s.githubClient.Search(ctx, q); err != nil {
		return nil, failure.Wrap(err)
	}
	return &Response{}, nil
}

func (req *Request) validate() error {
	return nil
}

func (req *Request) buildQuery() (string, error) {
	return "", fmt.Errorf("TBD")
}
