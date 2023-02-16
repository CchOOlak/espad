package main

import (
	"context"
	"espad/internal/core/service/urlsrv"
	"espad/internal/handlers/urlhdl"
	"espad/internal/repositories/urlrepo"
	"espad/pkg/hash"
	"espad/pkg/logging"
	"espad/pkg/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func main() {
	logging.Setup()

	// setup terminate channels
	ctx, cancel := context.WithCancel(context.Background())
	terminateNotify(ctx, cancel)

	// initialize services and repositories
	urlRepository := urlrepo.NewMemstorage()
	urlService := urlsrv.New(urlRepository, hash.New())
	urlHandler := urlhdl.NewHTTPHandler(urlService)

	// setup http server
	log.Info().Msg("starting http server on http://localhost:8585")
	router := server.InitRouter()
	
	router.POST("/add", urlHandler.Create)
	router.GET("/c/:path", urlHandler.Get)
	router.GET("/r/:path", urlHandler.Redirect)

	server.Start(ctx, router)
}

func terminateNotify(ctx context.Context, terminate context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	go handleSignals(signals, ctx, terminate)
}

func handleSignals(signals <-chan os.Signal, ctx context.Context, terminate context.CancelFunc) {
	<-signals
	shutdown(ctx)
	terminate()
}

func shutdown(ctx context.Context) {
	log.Info().Msg("server shutting down ...")
}
