package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/wire"
	"github.com/yckao/drone-convert-advanced/cmd/drone-convert-advanced/config"
	"github.com/yckao/drone-convert-advanced/core"
	"github.com/yckao/drone-convert-advanced/handler/convert"
	"github.com/yckao/drone-convert-advanced/handler/health"
	"github.com/yckao/drone-convert-advanced/server"
)

type (
	healthzHandler http.Handler
)

var serverSet = wire.NewSet(
	provideHealth,
	provideConvert,
	provideRouter,
	provideServer,
)

func provideRouter(healthz healthzHandler, convertz convert.Server) *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/healthz", healthz)
	r.Mount("/", convertz.Handler())
	return r
}

func provideHealth() healthzHandler {
	v := health.New()
	return healthzHandler(v)
}

func provideConvert(
	commits core.CommitService,
	repos core.RepositoryService,
	convertz core.ConvertService,
	config config.Config,
) convert.Server {
	return convert.New(commits, repos, convertz, config.Spec.Secret)
}

func provideServer(handler *chi.Mux, config config.Config) *server.Server {
	return &server.Server{
		Addr:    config.Spec.Bind,
		Handler: handler,
	}
}
