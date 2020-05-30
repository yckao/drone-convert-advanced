package runtime

import (
	"github.com/yckao/drone-convert-advanced/core"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

func fromChanged(v []*core.Change) *starlark.List {
	values := []starlark.Value{}
	for _, change := range v {
		values = append(values, starlarkstruct.FromStringDict(starlark.String("change"), starlark.StringDict{
			"path":    starlark.String(change.Path),
			"added":   starlark.Bool(change.Added),
			"renamed": starlark.Bool(change.Renamed),
			"deleted": starlark.Bool(change.Deleted),
		}))
	}
	return starlark.NewList(values)
}

func fromBuild(v core.Build) starlark.StringDict {
	return starlark.StringDict{
		"event":         starlark.String(v.Event),
		"action":        starlark.String(v.Action),
		"cron":          starlark.String(v.Cron),
		"environment":   starlark.String(v.Deploy),
		"link":          starlark.String(v.Link),
		"branch":        starlark.String(v.Target),
		"source":        starlark.String(v.Source),
		"before":        starlark.String(v.Before),
		"after":         starlark.String(v.After),
		"target":        starlark.String(v.Target),
		"ref":           starlark.String(v.Ref),
		"commit":        starlark.String(v.After),
		"title":         starlark.String(v.Title),
		"message":       starlark.String(v.Message),
		"source_repo":   starlark.String(v.Fork),
		"author_login":  starlark.String(v.Author),
		"author_name":   starlark.String(v.AuthorName),
		"author_email":  starlark.String(v.AuthorEmail),
		"author_avatar": starlark.String(v.AuthorAvatar),
		"sender":        starlark.String(v.Sender),
		"params":        fromMap(v.Params),
	}
}

func fromRepo(v core.Repository) starlark.StringDict {
	return starlark.StringDict{
		"uid":                  starlark.String(v.UID),
		"name":                 starlark.String(v.Name),
		"namespace":            starlark.String(v.Namespace),
		"slug":                 starlark.String(v.Slug),
		"git_http_url":         starlark.String(v.HTTPURL),
		"git_ssh_url":          starlark.String(v.SSHURL),
		"link":                 starlark.String(v.Link),
		"branch":               starlark.String(v.Branch),
		"config":               starlark.String(v.Config),
		"private":              starlark.Bool(v.Private),
		"visibility":           starlark.String(v.Visibility),
		"active":               starlark.Bool(v.Active),
		"trusted":              starlark.Bool(v.Trusted),
		"protected":            starlark.Bool(v.Protected),
		"ignore_forks":         starlark.Bool(v.IgnoreForks),
		"ignore_pull_requests": starlark.Bool(v.IgnorePulls),
	}
}

func fromMap(m map[string]string) *starlark.Dict {
	dict := new(starlark.Dict)
	for k, v := range m {
		dict.SetKey(
			starlark.String(k),
			starlark.String(v),
		)
	}
	return dict
}
