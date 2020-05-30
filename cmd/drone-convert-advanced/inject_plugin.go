package main

import (
	"github.com/drone/go-scm/scm"
	"github.com/google/wire"
	"github.com/yckao/drone-convert-advanced/core"
	"github.com/yckao/drone-convert-advanced/plugin/converter"
)

var pluginSet = wire.NewSet(
	provideStarlarkPlugin,
)

func provideStarlarkPlugin(client *scm.Client, contents core.FileService, commits core.CommitService) core.ConvertService {
	return converter.Combine(
		converter.Starlark(client, contents, commits),
	)
}
