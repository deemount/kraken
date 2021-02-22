// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package controllers ...
package controllers

import (
	"net/http"

	"github.com/deemount/kraken/api/responses"
)

// Home Controller
// @Summary Get home
// @Description Homepage
// @ID home
// @Accept json
// @Produce json
// @Header 200 {string} Token "ok"
// @Failure 404 {object} utils.HTTPError404
// @Router /home [get]
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Kraken API v1")
}
