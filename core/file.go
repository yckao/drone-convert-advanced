package core

import "context"

type (
	File struct {
		Data []byte
		Hash []byte
	}

	FileArgs struct {
		Commit string
		Ref    string
	}

	FileService interface {
		Find(ctx context.Context, repo, commit, ref, path string) (*File, error)
	}
)
