// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package controllers ...
package controllers

import (
	"net/http"

	"github.com/deemount/kraken/api/constants"
)

// initializeRoutes is a method
func (server *Server) initializeRoutes() error {

	var err error

	err = server.App.V1.HandleFunc(constants.HOMEURI, server.Home).Methods(http.MethodGet).GetError()
	if err != nil {
		return err
	}

	err = server.App.V1.HandleFunc(constants.BALANCEURI, server.GetBalance).Methods(http.MethodGet).GetError()
	if err != nil {
		return err
	}

	err = server.App.V1.HandleFunc(constants.TRADEBALANCEURI, server.GetTradeBalance).Methods(http.MethodGet).GetError()
	if err != nil {
		return err
	}

	err = server.App.V1.HandleFunc(constants.LEDGERURI, server.GetLedger).Methods(http.MethodGet).GetError()
	if err != nil {
		return err
	}

	return nil

}
