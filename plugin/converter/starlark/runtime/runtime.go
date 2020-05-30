package runtime

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/drone/go-scm/scm"
	"github.com/sirupsen/logrus"
	"github.com/yckao/drone-convert-advanced/core"
	"github.com/yckao/drone-convert-advanced/plugin/converter/starlark/loader"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

const (
	separator = "---"
	newline   = "\n"
)

const limit = 1000000

var (
	ErrMainMissing = errors.New("starlark: missing main function")
	ErrMainInvalid = errors.New("starlark: main must be a function")
	ErrMainReturn  = errors.New("starlark: main returns an invalid type")
	ErrMaximumSize = errors.New("starlark: maximum file size exceeded")
)

type Runtime struct {
	client  *scm.Client
	files   core.FileService
	commits core.CommitService
	repo    core.Repository
	build   core.Build
	config  core.Config
}

func New(client *scm.Client, files core.FileService, commits core.CommitService, repo core.Repository, build core.Build, config core.Config) *Runtime {
	return &Runtime{
		client:  client,
		files:   files,
		commits: commits,
		repo:    repo,
		build:   build,
		config:  config,
	}
}

func (r *Runtime) Run(ctx context.Context) (*core.Config, error) {
	loader := loader.New(r.client, r.files, r.repo, r.build, 30)
	thread := &starlark.Thread{
		Name: "drone",
		Load: loader.Load,
		Print: func(_ *starlark.Thread, msg string) {
			logrus.WithFields(logrus.Fields{
				"namespace": r.repo.Namespace,
				"name":      r.repo.Name,
			}).Traceln(msg)
		},
	}
	globals, err := starlark.ExecFile(thread, r.config.Data, []byte(r.config.Data), nil)
	if err != nil {
		return nil, err
	}

	mainVal, ok := globals["main"]
	if !ok {
		return nil, ErrMainMissing
	}
	main, ok := mainVal.(starlark.Callable)
	if !ok {
		return nil, ErrMainInvalid
	}

	args := r.createArgs()
	mainVal, err = starlark.Call(thread, main, args, nil)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	switch v := mainVal.(type) {
	case *starlark.List:
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			buf.WriteString(separator)
			buf.WriteString(newline)
			if err := write(buf, item); err != nil {
				return nil, err
			}
			buf.WriteString(newline)
		}
	case *starlark.Dict:
		if err := write(buf, v); err != nil {
			return nil, err
		}
	default:
		return nil, ErrMainReturn
	}

	if b := buf.Bytes(); len(b) > limit {
		return nil, ErrMaximumSize
	}

	return &core.Config{
		Data: buf.String(),
	}, nil
}

func (r *Runtime) createArgs() []starlark.Value {
	repoBranch := r.repo.Branch
	commitBranch := r.build.Source
	before := r.build.Before
	after := r.build.After

	if commitBranch != repoBranch {
		before = repoBranch
	}

	if before == "0000000000000000000000000000000000000000" || before == "" {
		before = fmt.Sprintf("%s~1", after)
	}

	changed, err := r.commits.CompareCommits(context.Background(), r.repo.Slug, before, after)
	if err != nil {
		changed = []*core.Change{}
		logrus.Debugln("starlark: failed listing changes in commit")
	}

	return []starlark.Value{
		starlarkstruct.FromStringDict(starlark.String("context"), starlark.StringDict{
			"repo":    starlarkstruct.FromStringDict(starlark.String("repo"), fromRepo(r.repo)),
			"build":   starlarkstruct.FromStringDict(starlark.String("build"), fromBuild(r.build)),
			"changes": fromChanged(changed),
		}),
	}
}
