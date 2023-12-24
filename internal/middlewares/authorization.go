package middlewares

import (
	"fmt"
	"net/http"

	"go-lessons/api"
	"go-lessons/internal/tools"

	"github.com/go-chi/chi/v5"
)

var errorUnAuthorized = fmt.Errorf("invalid token")

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = chi.URLParam(r, "username")
		var token string = r.Header.Get("Authorization")

		if token == "" || username == "" {
			api.SuccessResponse(w, nil, errorUnAuthorized.Error(), http.StatusUnauthorized, "error")
			return
		}

		var database *tools.DatabaseInterface = tools.NewDatabase()
		var loginDetails *tools.UserDetails = (*database).GetUserDetails(username)

		if loginDetails == nil {
			api.SuccessResponse(w, nil, errorUnAuthorized.Error(), http.StatusUnauthorized, "error")
			return
		}

		var tokenDB string = (*loginDetails).AuthToken
		if token != tokenDB {
			api.SuccessResponse(w, nil, errorUnAuthorized.Error(), http.StatusUnauthorized, "error")
			return
		}

		next.ServeHTTP(w, r)
	})
}
