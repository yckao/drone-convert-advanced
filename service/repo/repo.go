package repo

import (
	"context"

	"github.com/drone/go-scm/scm"
	"github.com/yckao/drone-convert-advanced/core"
)

type service struct {
	client *scm.Client
}

func New(client *scm.Client) core.RepositoryService {
	return &service{
		client: client,
	}
}

func (s *service) List(ctx context.Context) ([]*core.Repository, error) {
	repos := []*core.Repository{}
	opts := scm.ListOptions{Size: 100}
	for {
		result, meta, err := s.client.Repositories.List(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, src := range result {
			repos = append(repos, convertRepository(src))
		}
		opts.Page = meta.Page.Next
		opts.URL = meta.Page.NextURL

		if opts.Page == 0 && opts.URL == "" {
			break
		}
	}
	return repos, nil
}

func (s *service) Find(ctx context.Context, repo string) (*core.Repository, error) {
	result, _, err := s.client.Repositories.Find(ctx, repo)
	if err != nil {
		return nil, err
	}
	return convertRepository(result), nil
}
