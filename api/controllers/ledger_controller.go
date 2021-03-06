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
		log.Fatal(err)
	}

	responses.JSON(w, http.StatusOK, ledger)

}
