package core

import "context"

type (
	Commit struct {
		Sha       string
		Ref       string
		Message   string
		Author    *Committer
		Committer *Committer
		Link      string
	}

	Committer struct {
		Name   string
		Email  string
		Date   int64
		Login  string
		Avatar string
	}

	Change struct {
		Path    string
		Added   bool
		Renamed bool
		Deleted bool
	}

	CommitService interface {
		Find(ctx context.Context, repo, sha string) (*Commit, error)
		FindRef(ctx context.Context, repo, ref string) (*Commit, error)
		ListChanges(ctx context.Context, repo, sha string) ([]*Change, error)
		CompareCommits(ctx context.Context, repo, source, target string) ([]*Change, error)
	}
)
