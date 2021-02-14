package controllers

import (
	"net/http"

	"github.com/deemount/kraken/api/constants"
)

// initializeRoutes is a method
func (server *Server) initializeRoutes() {

	// uri
	home := constants.HOMEURI
	balance := constants.BALANCEURI

	//**************** Home Route

	server.App.V1.HandleFunc(home, server.Home).Methods(http.MethodGet)

	//**************** Balance Routes

	// single request
	server.App.V1.HandleFunc(balance, server.GetBalance).Methods(http.MethodGet)

}
