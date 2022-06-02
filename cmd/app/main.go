// Main is the application entrypoint
// Here we set up and start components including in HTTP Server and the Worker Pool
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/vitorsalgado/gopin/internal"
	"github.com/vitorsalgado/gopin/internal/domain"
	"github.com/vitorsalgado/gopin/internal/util/config"
	"github.com/vitorsalgado/gopin/internal/util/http/middlewares"
	"github.com/vitorsalgado/gopin/internal/util/observability"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/vitorsalgado/gopin/internal/util/db"
	"github.com/vitorsalgado/gopin/internal/util/worker"
)

const (
	applicationName = "GoPin"
	// ANSI Shadow
	banner = `
 ██████╗  ██████╗ ██████╗ ██╗███╗   ██╗
██╔════╝ ██╔═══██╗██╔══██╗██║████╗  ██║
██║  ███╗██║   ██║██████╔╝██║██╔██╗ ██║
██║   ██║██║   ██║██╔═══╝ ██║██║╚██╗██║
╚██████╔╝╚██████╔╝██║     ██║██║ ╚████║
 ╚═════╝  ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═══╝

%v
`
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	configurations := config.Load()
	database, _ := db.ConnectToMySQL(configurations)
	server, router := gopin.Server(configurations)

	fmt.Printf(banner, applicationName)
	fmt.Printf("Server starting on Port: %v\n", configurations.Port)
	fmt.Printf("Max Workers: %v\n\n", configurations.MaxWorkers)

	dispatcher := worker.NewDispatcher(configurations.MaxWorkers)
	dispatcher.Run()

	observability.ConfigureHealthCheck(router)
	gopin.Routes(router, dispatcher, domain.NewLocationRepository(database))
	router.ApplyRoutesTo(server)

	ext := make(chan os.Signal, 1)
	signal.Notify(ext, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	srv := http.Server{Addr: ":8080", Handler: middlewares.Recovery(server)}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		} else {
			log.Info().Msg("application stopped gracefully")
		}
	}()

	<-ext

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Panic().Err(err).Msgf("server shutdown failed. reason %s", err)
	}

	log.Info().Msg("server shutdown completed")
}
