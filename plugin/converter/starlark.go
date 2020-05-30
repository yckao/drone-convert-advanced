package converter

import (
	"context"
	"strings"

	"github.com/drone/go-scm/scm"
	"github.com/yckao/drone-convert-advanced/core"
	"github.com/yckao/drone-convert-advanced/plugin/converter/starlark/runtime"
)

func Starlark(client *scm.Client, files core.FileService, commits core.CommitService) core.ConvertService {
	return &starlarkPlugin{
		client:  client,
		files:   files,
		commits: commits,
	}
}

type starlarkPlugin struct {
	client  *scm.Client
	files   core.FileService
	commits core.CommitService
}

func (p *starlarkPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	switch {
	case strings.HasSuffix(req.Repo.Config, ".bzl"):
	case strings.HasSuffix(req.Repo.Config, ".star"):
	case strings.HasSuffix(req.Repo.Config, ".starlark"):
	default:
		return nil, nil
	}

	runtime := runtime.New(p.client, p.files, p.commits, *req.Repo, *req.Build, *req.Config)
	config, err := runtime.Run(ctx)

	if err != nil {
		return nil, err
	}

	return config, nil
}
