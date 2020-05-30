package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/drone/signal"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/yckao/drone-convert-advanced/cmd/drone-convert-advanced/config"
	"github.com/yckao/drone-convert-advanced/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	var envfile string
	flag.StringVar(&envfile, "env-file", ".env", "Read in a file of environment variables")
	flag.Parse()

	godotenv.Load(envfile)
	config, err := config.Environ()
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: invalid configuration")
	}

	initLogging(config)

	ctx := signal.WithContext(
		context.Background(),
	)

	if logrus.IsLevelEnabled(logrus.TraceLevel) {
		fmt.Println(config.String())
	}

	app, err := InitializeApplication(config)
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: cannot initialize server")
	}

	g := errgroup.Group{}
	g.Go(func() error {
		logrus.WithFields(
			logrus.Fields{
				"bind": config.Spec.Bind,
			},
		).Infoln("starting the http server")
		return app.server.ListenAndServe(ctx)
	})

	if err := g.Wait(); err != nil {
		logrus.WithError(err).Fatalln("program terminated")
	}
}

func initLogging(c config.Config) {
	if c.Logging.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if c.Logging.Trace {
		logrus.SetLevel(logrus.TraceLevel)
	}
	if c.Logging.Text {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   c.Logging.Color,
			DisableColors: !c.Logging.Color,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: c.Logging.Pretty,
		})
	}
}

// application is the main struct for the Drone server.
type application struct {
	server *server.Server
}

// newApplication creates a new application struct.
func newApplication(server *server.Server) application {
	return application{
		server: server,
	}
}
