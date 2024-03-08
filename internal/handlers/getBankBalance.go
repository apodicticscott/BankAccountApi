package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/apodicticscott/BankAccountApi/api"
	"github.com/apodicticscott/BankAccountApi/internal/tools"
	schema "github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetBankBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.BalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.BankDetails
	tokenDetails = (*database).GetUserBankDetails(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.BalanceResponse{
		Balance: (*tokenDetails).Cash,
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
