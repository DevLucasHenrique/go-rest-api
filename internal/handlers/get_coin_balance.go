package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DevLucasHenrique/go-rest-api/api"
	"github.com/DevLucasHenrique/go-rest-api/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params api.CoinBalanceParams
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&params, r.URL.Query()); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	database, err := tools.NewDatabase()
	if err != nil || database == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	coinDetails := (*database).GetUserCoins(params.Username)
	if coinDetails == nil {
		api.RequestErrorHandler(w, errors.New("user not found"))
		return
	}

	response := api.CoinBalanceResponse{
		StatusCode: http.StatusOK,
		Balance:    coinDetails.Coins,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
