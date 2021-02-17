package controllers

import (
	"log"
	"net/http"

	"github.com/deemount/kraken/api/constants"
	"github.com/deemount/kraken/api/repositories"
	"github.com/deemount/kraken/api/responses"
)

// GetTradeBalance Controller
// @Summary Get trade balance
// @Description Calling the kraken api and get the trade balance
// @ID get-trade-balance
// @Accept json
// @Produce json
// @Success 200 {object} models.TradeBalance
// @Header 200 {string} Token "ok"
// @Failure 404 {object} http.
// @Router /balance [get]
func (server *Server) GetTradeBalance(w http.ResponseWriter, r *http.Request) {

	service := repositories.NewTradeBalanceService(
		server.App.API.Version,
		server.App.Kraken.URL,
		constants.BALANCEURI,
		server.App.Kraken.UserAgent,
		server.App.Kraken.Key,
		server.App.Kraken.Secret)

	repository := repositories.TradeBalanceRepository(service)

	args := map[string]string{
		"aclass": "currency",
		"asset":  "XETH",
	}

	tradeBalance, err := repository.GetTradeBalance(args)
	if err != nil {
		log.Print(err)
	}

	responses.JSON(w, http.StatusOK, tradeBalance)

}
