package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/titobzzz/goapi/internal/tools"
	"github.com/titobzzz/internal/api"
)

var UnAuthorizedError = errors.New("Invalid username or token..")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err errors

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorhnadler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginData(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorhnadler(w, UnAuthorizedError)
			return
		}

		next.ServceHTTP(w, r)

	})
}
