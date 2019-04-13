package main

import (
	"fmt"
	"github.com/Yurovskikh/news/api/config"
	"github.com/Yurovskikh/news/api/pkg/service"
	"github.com/Yurovskikh/news/api/pkg/transport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	newsService := service.NewNewsService(cfg)
	endpoints := transport.MakeEndpoints(newsService, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		srv := &http.Server{
			Handler:           transport.NewHttpService(endpoints, logger),
			Addr:              fmt.Sprintf(":%d", cfg.AppPort),
			WriteTimeout:      30 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			MaxHeaderBytes:    1 << 20,
		}
		errs <- srv.ListenAndServe()
	}()
	level.Error(logger).Log("exit", <-errs)
	defer func() {
		level.Info(logger).Log("msg", "service ended")
	}()
}
