//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/yckao/drone-convert-advanced/cmd/drone-convert-advanced/config"
)

func InitializeApplication(config config.Config) (application, error) {
	wire.Build(
		clientSet,
		pluginSet,
		serverSet,
		serviceSet,
		newApplication,
	)
	return application{}, nil
}
