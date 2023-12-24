package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type CoinBalanceParams struct {
	Username string `json:"username"`
}

type UpdateCoinBalanceBody struct {
	Coins int64 `json:"coins" validate:"required"`
}

type CoinBalanceResponse struct {
	Balance int64 `json:"balance"`
}

type NewUserBody struct {
	Username string `json:"username" validate:"required"`
	Token    string `json:"token" validate:"required"`
	Coins    int64  `json:"coins" validate:"required"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta"`
}

func SuccessResponse(w http.ResponseWriter, meta interface{}, message string, code int, status string) {
	resp := Response{
		Status:  status,
		Message: message,
		Meta:    meta,
	}

	if status == "error" {
		log.Error(message)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				response := Response{
					Status:  "error",
					Message: fmt.Sprintf("%v", r), // Convert r to string using fmt.Sprintf
					Meta:    nil,
				}
				json.NewEncoder(w).Encode(response)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
