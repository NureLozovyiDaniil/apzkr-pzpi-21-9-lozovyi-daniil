package main

import (
	"auth/domain"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"log"
	"os"
)

func main() {
	dbURI := "mongodb://localhost:27017/"

	repo, err := domain.NewMongoRepo(dbURI)
	if err != nil {
		log.Fatal(err)
	}

	var (
		logger  = newLogger()
		svc     = domain.NewService(repo, logger)
		handler = domain.NewHandler(svc)
		router  = handler.InitRoutes()
	)
	defer level.Info(logger).Log("msg", "service ended")

	if err := router.Start(":5000"); err != nil {
		log.Fatal(err)
	}
}

func newLogger() kitlog.Logger {
	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stderr)
		logger = kitlog.NewSyncLogger(logger)
		logger = kitlog.With(logger,
			"service", "auth",
			"time:", kitlog.DefaultTimestampUTC,
			"caller", kitlog.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")

	return logger
}
