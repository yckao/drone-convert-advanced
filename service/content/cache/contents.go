package cache

import (
	"context"
	"fmt"

	lru "github.com/hashicorp/golang-lru"
	"github.com/yckao/drone-convert-advanced/core"
)

const contentKey = "%s/%s/%s"

func Contents(base core.FileService) core.FileService {
	cache, _ := lru.New(25)
	return &service{
		service: base,
		cache:   cache,
	}
}

type service struct {
	cache   *lru.Cache
	service core.FileService
}

func (s *service) Find(ctx context.Context, repo, commit, ref, path string) (*core.File, error) {
	key := fmt.Sprintf(contentKey, repo, commit, path)
	cached, ok := s.cache.Get(key)
	if ok {
		return cached.(*core.File), nil
	}
	file, err := s.service.Find(ctx, repo, commit, ref, path)
	if err != nil {
		return nil, err
	}
	s.cache.Add(key, file)
	return file, nil
}
