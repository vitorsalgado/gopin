// Main is the application entrypoint
// Here we set up and start components including in HTTP Server and the Worker Pool
package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/vitorsalgado/go-location-management/internal/locations"
	"github.com/vitorsalgado/go-location-management/internal/server/rest"
	"github.com/vitorsalgado/go-location-management/internal/server/rest/middlewares"
	"github.com/vitorsalgado/go-location-management/internal/utils/config"
	"github.com/vitorsalgado/go-location-management/internal/utils/database"
	"github.com/vitorsalgado/go-location-management/internal/utils/panicif"
	"github.com/vitorsalgado/go-location-management/internal/utils/worker"
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
	db := database.Connect(configurations)
	server, router := rest.Server(configurations)

	dispatcher := worker.NewDispatcher(configurations.MaxWorkers)
	dispatcher.Run()

	rest.RegisterRoutes(router)
	locations.RegisterRoutes(router, dispatcher, locations.NewRepository(db))

	router.ApplyRoutesTo(server)

	fmt.Print(fmt.Sprintf(banner, applicationName))
	fmt.Println(fmt.Sprintf("Server starting on Port: %v", configurations.Port))
	fmt.Println(fmt.Sprintf("Max Workers: %v", configurations.MaxWorkers))
	fmt.Println("::")

	panicif.Err(http.ListenAndServe(":8080", middlewares.Recovery(server)))
}
