package core

import "context"

const (
	VisibilityPublic   = "public"
	VisibilityPrivate  = "private"
	VisibilityInternal = "internal"
)

type (
	Repository struct {
		ID          int64  `json:"id"`
		UID         string `json:"uid"`
		UserID      int64  `json:"user_id"`
		Namespace   string `json:"namespace"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		SCM         string `json:"scm"`
		HTTPURL     string `json:"git_http_url"`
		SSHURL      string `json:"git_ssh_url"`
		Link        string `json:"link"`
		Branch      string `json:"default_branch"`
		Private     bool   `json:"private"`
		Visibility  string `json:"visibility"`
		Active      bool   `json:"active"`
		Config      string `json:"config_path"`
		Trusted     bool   `json:"trusted"`
		Protected   bool   `json:"protected"`
		IgnoreForks bool   `json:"ignore_forks"`
		IgnorePulls bool   `json:"ignore_pull_requests"`
		CancelPulls bool   `json:"auto_cancel_pull_requests"`
		CancelPush  bool   `json:"auto_cancel_pushes"`
		Timeout     int64  `json:"timeout"`
		Counter     int64  `json:"counter"`
		Synced      int64  `json:"synced"`
		Created     int64  `json:"created"`
		Updated     int64  `json:"updated"`
		Version     int64  `json:"version"`
		Signer      string `json:"signer,omitempty"`
		Secret      string `json:"secret,omitempty"`
		Build       Build  `json:"build,omitempty"`
	}

	RepositoryService interface {
		List(ctx context.Context) ([]*Repository, error)
		Find(ctx context.Context, repo string) (*Repository, error)
	}
)
