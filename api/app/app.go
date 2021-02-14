package app

import (
	"github.com/gorilla/mux"

	"github.com/deemount/kraken/api/config"
	"github.com/deemount/kraken/api/config/driver"
)

// App ...
type App struct {
	// refer
	DB     driver.DataService
	Routes config.Routes

	// pointer
	API     *config.API
	Kraken  *config.Kraken
	Options *config.Options
	Swagger *config.Swagger

	// legal
	Router *mux.Router
	V1     *mux.Router
}
