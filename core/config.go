package core

import "context"

type (
	Config struct {
		Data string `json:"data"`
		Kind string `json:"kind"`
	}

	ConfigArgs struct {
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`
		Config *Config     `json:"config,omitempty"`
	}

	ConfigService interface {
		Find(context.Context, *ConfigArgs) (*Config, error)
	}
)
