package middleware

import (
	"errors"
	"net/http"

	"github.com/DevLucasHenrique/go-rest-api/api"
	"github.com/DevLucasHenrique/go-rest-api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		database, err := tools.NewDatabase()
		if err != nil || database == nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		loginDetails := (*database).GetUserLoginDetails(username)
		if loginDetails == nil || token != loginDetails.AuthToken {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
