package main

import (
	"github.com/drone/go-scm/scm"
	"github.com/google/wire"
	"github.com/yckao/drone-convert-advanced/cmd/drone-convert-advanced/config"
	"github.com/yckao/drone-convert-advanced/core"
	"github.com/yckao/drone-convert-advanced/service/commit"
	contents "github.com/yckao/drone-convert-advanced/service/content"
	"github.com/yckao/drone-convert-advanced/service/content/cache"
	"github.com/yckao/drone-convert-advanced/service/repo"
)

var serviceSet = wire.NewSet(
	commit.New,
	provideRepositoryService,
	provideContentService,
)

func provideContentService(client *scm.Client) core.FileService {
	return cache.Contents(
		contents.New(client),
	)
}

func provideRepositoryService(client *scm.Client, config config.Config) core.RepositoryService {
	return repo.New(
		client,
	)
}
