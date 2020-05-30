package config

import (
	"strings"

	"github.com/ghodss/yaml"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type (
	Config struct {
		Spec    Spec
		Logging Logging

		Bitbucket Bitbucket
		Gitea     Gitea
		Github    Github
		GitLab    GitLab
		Gogs      Gogs
	}

	Logging struct {
		Debug  bool
		Trace  bool `envconfig:"DRONE_LOGS_TRACE"`
		Color  bool `envconfig:"DRONE_LOGS_COLOR"`
		Pretty bool `envconfig:"DRONE_LOGS_PRETTY"`
		Text   bool `envconfig:"DRONE_LOGS_TEXT"`
	}

	Spec struct {
		Bind   string `envconfig:"DRONE_BIND"`
		Debug  bool   `envconfig:"DRONE_DEBUG"`
		Secret string `envconfig:"DRONE_SECRET"`
	}

	Bitbucket struct {
		Username    string `envconfig:"DRONE_BITBUCKET_USERNAME"`
		AppPassword string `envconfig:"DRONE_BITBUCKET_APP_PASSWORD"`
		Debug       bool   `envconfig:"DRONE_BITBUCKET_DEBUG"`
		SkipVerify  bool   `envconfig:"DRONE_BITBUCKET_SKIP_VERIFY"`
	}
	Gitea struct {
		Server     string `envconfig:"DRONE_GITEA_SERVER"`
		Token      string `envconfig:"DRONE_GITEA_TOKEN"`
		Debug      bool   `envconfig:"DRONE_GITEA_DEBUG"`
		SkipVerify bool   `envconfig:"DRONE_GITEA_SKIP_VERIFY"`
	}
	Github struct {
		Server     string `envconfig:"DRONE_GITHUB_SERVER" default:"https://github.com"`
		APIServer  string `envconfig:"DRONE_GITHUB_API_SERVER"`
		Token      string `envconfig:"DRONE_GITHUB_TOKEN"`
		RateLimit  int    `envconfig:"DRONE_GITHUB_USER_RATELIMIT"`
		Debug      bool   `envconfig:"DRONE_GITHUB_DEBUG"`
		SkipVerify bool   `envconfig:"DRONE_GITHUB_SKIP_VERIFY"`
	}
	GitLab struct {
		Server     string `envconfig:"DRONE_GITLAB_SERVER" default:"https://gitlab.com"`
		Token      string `envconfig:"DRONE_GITLAB_TOKEN"`
		Debug      bool   `envconfig:"DRONE_GITLAB_DEBUG"`
		SkipVerify bool   `envconfig:"DRONE_GITLAB_SKIP_VERIFY"`
	}
	Gogs struct {
		Server     string `envconfig:"DRONE_GOGS_SERVER"`
		Token      string `envconfig:"DRONE_GOGS_TOKEN"`
		Debug      bool   `envconfig:"DRONE_GOGS_DEBUG"`
		SkipVerify bool   `envconfig:"DRONE_GOGS_SKIP_VERIFY"`
	}
)

func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	configureGithub(&cfg)

	return cfg, err
}

func (c *Config) String() string {
	out, _ := yaml.Marshal(c)
	return string(out)
}

func (c *Config) IsGithub() bool {
	return c.Github.Token != ""
}

func (c *Config) IsGitHubEnterprise() bool {
	return c.IsGithub() && !strings.HasPrefix(c.Github.Server, "https://github.com")
}

func (c *Config) IsGitLab() bool {
	return c.GitLab.Token != ""
}

func (c *Config) IsGogs() bool {
	return c.Gogs.Server != ""
}

func (c *Config) IsGitea() bool {
	return c.Gitea.Server != ""
}

func (c *Config) IsBitbucket() bool {
	return c.Bitbucket.Username != ""
}

func configureGithub(c *Config) {
	if c.Github.APIServer != "" {
		return
	}
	if c.Github.Server == "https://github.com" {
		c.Github.APIServer = "https://api.github.com"
	} else {
		c.Github.Server = strings.TrimSuffix(c.Github.Server, "/") + "/api/v3"
	}
}

func defaultSettings(c *Config) {
	if c.Spec.Debug {
		c.Logging.Debug = c.Spec.Debug
	}

	if c.Spec.Secret == "" {
		logrus.Fatalln("missing secret key")
	}

	if c.Spec.Bind == "" {
		c.Spec.Bind = ":3000"
	}
}
