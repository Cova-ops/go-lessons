package account

import (
	"encoding/json"
	"net/http"

	"go-lessons/api"
	"go-lessons/internal/tools"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var database *tools.DatabaseInterface = tools.NewDatabase()
	var users []string = (*database).GetUsers()

	api.SuccessResponse(w, users, "Users Accepted", http.StatusOK, "success")
}

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var username string = chi.URLParam(r, "username")

	var database *tools.DatabaseInterface = tools.NewDatabase()
	var tokenDetails *tools.UserDetails = (*database).GetUserDetails(username)

	if tokenDetails == nil {
		api.SuccessResponse(w, nil, "User not founded!", http.StatusBadRequest, "error")
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
	}

	api.SuccessResponse(w, response, "User Accepted", http.StatusOK, "success")
}

func UpdateCoinBalance(w http.ResponseWriter, r *http.Request) {
	var username string = chi.URLParam(r, "username")

	var body api.UpdateCoinBalanceBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		api.SuccessResponse(w, nil, err.Error(), http.StatusBadRequest, "error")
		return
	}

	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		api.SuccessResponse(w, nil, err.Error(), http.StatusBadRequest, "error")
		return
	}

	var database *tools.DatabaseInterface = tools.NewDatabase()
	var success bool = (*database).UpdateUserCoins(username, body.Coins)

	if !success {
		api.SuccessResponse(w, nil, "User not founded", http.StatusBadRequest, "error")
		return
	}

	api.SuccessResponse(w, tools.CoinDetails{Username: username, Coins: body.Coins}, "User Updated", http.StatusAccepted, "success")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	var body api.NewUserBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		api.SuccessResponse(w, nil, err.Error(), http.StatusBadRequest, "error")
		return
	}

	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		api.SuccessResponse(w, nil, err.Error(), http.StatusBadRequest, "error")
		return
	}

	var database *tools.DatabaseInterface = tools.NewDatabase()
	var success bool = (*database).NewUser(body.Username, body.Token, body.Coins)

	if !success {
		api.SuccessResponse(w, nil, "User already exists", http.StatusBadRequest, "error")
		return
	}

	api.SuccessResponse(w, body, "User Created", http.StatusCreated, "success")
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	var username string = chi.URLParam(r, "username")

	var database *tools.DatabaseInterface = tools.NewDatabase()
	var success bool = (*database).RemoveUser(username)

	if !success {
		api.SuccessResponse(w, nil, "User not founded", http.StatusBadRequest, "error")
		return
	}

	api.SuccessResponse(w, api.CoinBalanceParams{Username: username}, "User Removed", http.StatusAccepted, "success")
}
