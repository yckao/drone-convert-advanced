package core

import "context"

type (
	ConvertArgs struct {
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`
		Config *Config     `json:"config,omitempty"`
	}

	ConvertService interface {
		Convert(context.Context, *ConvertArgs) (*Config, error)
	}
)
