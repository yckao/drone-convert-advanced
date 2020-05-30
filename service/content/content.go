package contents

import (
	"context"
	"time"

	"github.com/drone/go-scm/scm"
	"github.com/yckao/drone-convert-advanced/core"
)

var attempts = 3
var wait = time.Second * 5

func New(client *scm.Client) core.FileService {
	return &service{
		client:   client,
		attempts: attempts,
		wait:     wait,
	}
}

type service struct {
	client   *scm.Client
	attempts int
	wait     time.Duration
}

func (s *service) Find(ctx context.Context, repo, commit, ref, path string) (*core.File, error) {
	// TODO: Figure out why gogs client have an workground in github.com/drone/drone/service/content
	content, err := s.findRetry(ctx, repo, path, commit)
	if err != nil {
		return nil, err
	}
	return &core.File{
		Data: content.Data,
		Hash: []byte{},
	}, nil
}

func (s *service) findRetry(ctx context.Context, repo, path, commit string) (content *scm.Content, err error) {
	for i := 0; i < s.attempts; i++ {
		content, _, err = s.client.Contents.Find(ctx, repo, path, commit)
		if err == nil {
			return
		}

		time.Sleep(s.wait)
	}
	return
}
