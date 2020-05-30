package loader

import (
	"context"
	"errors"
	"path"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/drone/go-scm/scm"
	"github.com/sirupsen/logrus"
	"github.com/yckao/drone-convert-advanced/core"
	"go.starlark.net/starlark"
)

var (
	ErrExternalDependencies = errors.New("starlark: currently not support external dependencies")
	ErrInvalidExtension     = errors.New("starlark: can't load files that are not starlark extensions")
	ErrExceedMaxFileCount   = errors.New("starlark: exceed max file count")
)

type Loader struct {
	client *scm.Client
	files  core.FileService
	repo   core.Repository
	build  core.Build

	cache map[string]string

	limit int

	count int
}

func New(client *scm.Client, files core.FileService, repo core.Repository, build core.Build, limit int) *Loader {
	return &Loader{
		client: client,
		files:  files,
		repo:   repo,
		build:  build,
		limit:  limit,
	}
}

func (l *Loader) Load(t *starlark.Thread, labelStr string) (starlark.StringDict, error) {
	pLabel, err := parseLabelStr(labelStr)
	if err != nil {
		logger := logrus.WithError(err)
		logger.Debugln("starlark: error while parsing extension label")
		return nil, err
	}

	path := pathFromLabel(pLabel)

	loaded, err := l.loadFromFile(t, path)
	if err != nil {
		logrus.WithError(err).Debugln("starlark: weeor while load from file")
		return nil, err
	}

	return loaded, nil
}

func (l *Loader) loadFromFile(t *starlark.Thread, path string) (starlark.StringDict, error) {
	if l.cache == nil {
		l.cache = map[string]string{}
	}

	if content, ok := l.cache[path]; ok {
		return starlark.ExecFile(t, path, content, nil)
	}

	defer func() {
		l.count++
	}()

	if l.limit > 0 && l.count > l.limit {
		return nil, ErrExceedMaxFileCount
	}

	ctx := context.Background()
	content, err := l.files.Find(ctx, l.repo.Slug, l.build.After, l.build.Ref, path)
	if err != nil {
		return nil, err
	}

	return starlark.ExecFile(t, path, content.Data, nil)
}

func parseLabelStr(labelStr string) (*label.Label, error) {
	parsed, err := label.Parse(labelStr)
	if err != nil {
		return nil, err
	}

	if parsed.Repo != "" {
		return nil, ErrExternalDependencies
	}
	if !isValidFilename(parsed.Name) {
		return nil, ErrInvalidExtension
	}

	return &parsed, nil
}

func pathFromLabel(pLabel *label.Label) string {
	if pLabel.Relative {
		return pLabel.Name
	}
	return path.Join(pLabel.Pkg, pLabel.Name)
}

func isValidFilename(name string) bool {
	switch {
	case strings.HasSuffix(name, ".star"):
		return true
	case strings.HasSuffix(name, ".starlark"):
		return true
	case strings.HasSuffix(name, ".bzl"):
		return true
	default:
		return false
	}
}
