package main

import (
	"aid/domain"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

func main() {
	logger := newLogger()

	dbURI := getEnv("MONGODB_URI", "mongodb://localhost:27017/")

	var (
		svc     = domain.NewSvc(dbURI)
		handler = domain.NewHandler(svc)
		router  = handler.InitRoutes()
	)

	level.Info(logger).Log("msg", "starting server", "addr", ":8080")
	if err := router.Start(":5001"); err != nil {
		level.Error(logger).Log("msg", "server stopped", "err", err)
	}
}

func newLogger() kitlog.Logger {
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	logger = kitlog.NewSyncLogger(logger)
	logger = kitlog.With(logger,
		"service", "aid",
		"time:", kitlog.DefaultTimestampUTC,
		"caller", kitlog.DefaultCaller,
	)
	return logger
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
