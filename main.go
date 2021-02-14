package main

import (
	"log"

	"github.com/deemount/kraken/api"
)

// @title Kraken REST API
// @version 0.1.0
// @description Fetches data from Kraken.com and stores it in database
// @termsOfService https://github.com/deemount/kraken/terms/index.html

// @contact.name API Support
// @contact.url https://github.com/deemount
// @contact.email salvatore.gonda@web.de

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8686
// @BasePath /kraken/v1
func main() {

	// assign error
	var err error

	// run application interface
	if err = api.Run(); err != nil {
		log.Fatalf("Kraken REST API Error %+s", err)
	}

}
