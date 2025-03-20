package handlers

import(
	"encoding/json"
	"net/http"

	"github.com/shivGam/goapi/api"
	"github.com/shivGam/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter,r *http.Request){
	params:= api.CoinBalanceParams{}
	decoder := schema.NewDecoder()
	var err error
	err = decoder.Decode(&params,r.URL.Query())

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

	tokenDetails := (*database).GetUserCoinDetails(params.Username)
	if tokenDetails ==nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	resp := api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}