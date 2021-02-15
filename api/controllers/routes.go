package controllers

import (
	"net/http"

	"github.com/deemount/kraken/api/constants"
)

// initializeRoutes is a method
func (server *Server) initializeRoutes() error {

	var err error

	// uri
	home := constants.HOMEURI
	balance := constants.BALANCEURI

	//**************** Home Route

	err = server.App.V1.HandleFunc(home, server.Home).Methods(http.MethodGet).GetError()
	if err != nil {
		return err
	}

	//**************** Balance Routes

	// single request
	err = server.App.V1.HandleFunc(balance, server.GetBalance).Methods(http.MethodGet).GetError()
	if err != nil {
		return err
	}

	return nil

}
