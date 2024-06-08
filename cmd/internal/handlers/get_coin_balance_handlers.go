package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/titobzzz/gopi/api"
	"github.com/titobzzz/gopi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder = schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Deode(&params, r.URL.Query().)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writeError
		return)
}

var database *tools.DatabaseInterface
database, err = tools.NewDatabase()
if err != nil {
    log.Error(err)
    api.InternalErrorHandler(w)
    return
}

var tokenDetails *tools.CoinDetails
database, err = tools.NewDatabase()
if err != nil {
    log.Error(err)
    api.InternalErrorHandler(w)
    return
}

var response = api.CoinBalanceResponse{

}