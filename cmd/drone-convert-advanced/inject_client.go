package main

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/bitbucket"
	"github.com/drone/go-scm/scm/driver/gitea"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/driver/gitlab"
	"github.com/drone/go-scm/scm/driver/gogs"
	"github.com/drone/go-scm/scm/transport"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/yckao/drone-convert-advanced/cmd/drone-convert-advanced/config"
)

var clientSet = wire.NewSet(
	provideClient,
)

func provideClient(config config.Config) *scm.Client {
	switch {
	case config.Bitbucket.Username != "":
		return provideBitbucketClient(config)
	case config.Github.Token != "":
		return provideGithubClient(config)
	case config.Gitea.Server != "":
		return provideGiteaClient(config)
	case config.GitLab.Token != "":
		return provideGitlabClient(config)
	case config.Gogs.Server != "":
		return provideGogsClient(config)
	}
	logrus.Fatalln("main: source code management system not configured")
	return nil
}

func provideBitbucketClient(config config.Config) *scm.Client {
	client := bitbucket.NewDefault()
	client.Client = &http.Client{
		Transport: &transport.BasicAuth{
			Username: config.Bitbucket.Username,
			Password: config.Bitbucket.AppPassword,
			Base:     defaultTransport(config.Bitbucket.SkipVerify),
		},
	}
	if config.Bitbucket.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	return client
}

func provideGithubClient(config config.Config) *scm.Client {
	client, err := github.New(config.Github.APIServer)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the GitHub client")
	}
	if config.Github.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &transport.BearerToken{
			Token: config.Github.Token,
			Base:  defaultTransport(config.Github.SkipVerify),
		},
	}
	return client
}

func provideGiteaClient(config config.Config) *scm.Client {
	client, err := gitea.New(config.Gitea.Server)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the Gitea client")
	}
	if config.Gitea.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &transport.BearerToken{
			Token: config.Gitea.Token,
			Base:  defaultTransport(config.Gitea.SkipVerify),
		},
	}
	return client
}

func provideGitlabClient(config config.Config) *scm.Client {
	logrus.WithField("server", config.GitLab.Server).
		WithField("skip_verify", config.GitLab.SkipVerify).
		Debugln("main: creating the GitLab client")
	client, err := gitlab.New(config.GitLab.Server)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the GitLab client")
	}
	if config.GitLab.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &transport.BearerToken{
			Token: config.GitLab.Token,
			Base:  defaultTransport(config.GitLab.SkipVerify),
		},
	}
	return client
}

func provideGogsClient(config config.Config) *scm.Client {
	logrus.WithField("server", config.Gogs.Server).
		WithField("skip_verify", config.Gogs.SkipVerify).
		Debugln("main: creating the Gogs client")

	client, err := gogs.New(config.Gogs.Server)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the Gogs client")
	}
	if config.Gogs.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &transport.BearerToken{
			Token: config.Gogs.Token,
			Base:  defaultTransport(config.Gogs.SkipVerify),
		},
	}
	return client
}

func defaultTransport(skipverify bool) http.RoundTripper {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipverify,
		},
	}
}
