package controllers

import (
	"log"
	"net/http"

	"github.com/deemount/kraken/api/constants"
	"github.com/deemount/kraken/api/repositories"
	"github.com/deemount/kraken/api/responses"
)

// GetBalance Controller
// @Summary Get balance
// @Description Calling the kraken api and get the balance
// @ID get-balance
// @Accept json
// @Produce json
// @Success 200 {object} models.Balance
// @Header 200 {string} Token "ok"
// @Failure 404 {object} http.
// @Router /balance [get]
func (server *Server) GetBalance(w http.ResponseWriter, r *http.Request) {

	service := repositories.NewBalanceService(
		server.App.API.Version,
		server.App.Kraken.URL,
		constants.BALANCEURI,
		server.App.Kraken.UserAgent,
		server.App.Kraken.Key,
		server.App.Kraken.Secret)

	repository := repositories.BalanceRepository(service)

	result, err := repository.GetBalance()
	if err != nil {
		log.Fatalf("controller: %s", err)
	}

	responses.JSON(w, http.StatusOK, result)

}
