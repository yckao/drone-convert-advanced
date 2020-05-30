package repo

import (
	"github.com/drone/go-scm/scm"
	"github.com/yckao/drone-convert-advanced/core"
)

func convertRepository(src *scm.Repository) *core.Repository {
	return &core.Repository{
		UID:        src.ID,
		Namespace:  src.Namespace,
		Name:       src.Name,
		Slug:       scm.Join(src.Namespace, src.Name),
		HTTPURL:    src.Clone,
		SSHURL:     src.CloneSSH,
		Link:       src.Link,
		Private:    src.Private,
		Visibility: convertVisibility(src),
		Branch:     src.Branch,
	}
}

func convertVisibility(src *scm.Repository) string {
	switch {
	case src.Private == true:
		return core.VisibilityPrivate
	default:
		return core.VisibilityPublic
	}
}
