package main

import (
	"fmt"
	"github.com/Yurovskikh/news/storage/config"
	"github.com/Yurovskikh/news/storage/pkg"
	"github.com/Yurovskikh/news/storage/pkg/repository"
	"github.com/Yurovskikh/news/storage/pkg/service"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := log.NewJSONLogger(os.Stdout)
	logger = log.NewSyncLogger(logger)
	logger = level.NewFilter(logger, level.AllowDebug())
	logger = log.With(
		logger,
		"time", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	level.Info(logger).Log("msg", "service starting")

	cfg := config.ReadConfig()

	db := repository.NewDB(cfg)
	newsService := service.NewNewsService(repository.NewNewsRepository(db))

	msgHandler := pkg.NewMsgHandler(cfg, newsService, logger)
	msgHandler.Start()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	level.Error(logger).Log("exit", <-errs)
	defer func() {
		level.Info(logger).Log("msg", "service ended")
	}()
}
