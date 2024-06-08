package api

import {
	"encoding/json"
	"net/http"
}

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int

	Balance int64
}

type CError struct {
	Code int

	Message sting
}

func writeError(w http.ResponseWriter, Message string, code int) {
	resp := Error{
		Code: code
		Message: Message
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}