// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package app ...
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
	API     config.API
	Kraken  config.Kraken
	Options config.Options
	Swagger config.Swagger

	// legal
	Router *mux.Router
	V1     *mux.Router
}
