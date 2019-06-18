package cli

import (
	"context"
	"fmt"
	"github.com/keizo042/hubq/config"
	"github.com/keizo042/hubq/di"
	"github.com/keizo042/hubq/search"
	"github.com/morikuni/failure"
	"github.com/urfave/cli"
)

func (com *Commands) SearchFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "keyword, k",
			Value: "",
		},
		cli.StringFlag{
			Name:  "execute, e",
			Value: "",
		},
	}
}

type searchCommand struct {
	searchService search.Search
}

func newSearchCommand(searchService search.Search) (*searchCommand, error) {
	return &searchCommand{
		searchService: searchService,
	}, nil
}

func (com *searchCommand) parse(c *cli.Context) (*search.Request, *search.Option, error) {
	res := &search.Request{
		RawQuery: c.String("execute"),
		Keyword:  c.String("keyword"),
	}
	return res, nil, nil
}

func (com *searchCommand) show(res *search.Response) error {
	fmt.Println("name\trepo\turl")
	for _, repo := range res.Repositories {
		fmt.Println(repo.CloneURL)
	}
	return nil
}

func (com *Commands) SearchCommand() cli.Command {
	return cli.Command{
		Name:   "search",
		Action: com.Search,
		Flags:  com.SearchFlags(),
	}
}

func (com *Commands) Search(c *cli.Context) error {
	ctx := context.Background()
	cfg, err := config.Parse(com.ConfigPath)
	if err != nil {
		return failure.Wrap(err)
	}

	container, err := di.NewContainer(cfg)
	if err != nil {
		return failure.Wrap(err)
	}

	githubClient, err := container.InjectGithub(ctx)
	if err != nil {
		return failure.Wrap(err)
	}

	searchService, err := search.New(githubClient)
	if err != nil {
		return failure.Wrap(err)
	}

	searchCommand, err := newSearchCommand(searchService)
	if err != nil {
		return failure.Wrap(err)
	}

	req, opt, err := searchCommand.parse(c)
	if err != nil {
		return failure.Wrap(err)
	}

	res, err := searchService.Search(ctx, req, opt)
	if err != nil {
		return failure.Wrap(err)
	}

	if err := searchCommand.show(res); err != nil {
		return failure.Wrap(err)
	}
	return nil
}
