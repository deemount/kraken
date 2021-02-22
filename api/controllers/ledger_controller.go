// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package controllers ...
package controllers

import (
	"log"
	"net/http"

	"github.com/deemount/kraken/api/constants"
	"github.com/deemount/kraken/api/repositories"
	"github.com/deemount/kraken/api/responses"
)

// GetLedger Controller
// @Summary Get ledger
// @Description Calling the kraken api and get the ledger
// @ID get-ledger
// @Accept json
// @Produce json
// @Success 200 {object} models.Ledger
// @Header 200 {string} Token "ok"
// @Failure 404 {object} http.
// @Router /ledger [get]
func (server *Server) GetLedger(w http.ResponseWriter, r *http.Request) {

	service := repositories.NewLedgerService(
		server.App.API.Version,
		server.App.Kraken.URL,
		constants.LEDGERURI,
		server.App.Kraken.UserAgent,
		server.App.Kraken.Key,
		server.App.Kraken.Secret)

	repository := repositories.LedgerRepository(service)

	args := map[string]string{
		"aclass": "currency",
		"asset":  "all",
		"type":   "all",
		"start":  "",
		"end":    "",
		"ofs":    "",
	}

	ledger, err := repository.GetLedger(args)
	if err != nil {
		log.Fatalf("controller: %s", err)
	}

	responses.JSON(w, http.StatusOK, ledger)

}
