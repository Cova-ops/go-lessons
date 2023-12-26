package account

import (
	"encoding/json"
	"net/http"

	"go-lessons/api"
	"go-lessons/internal/tools"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type ResponseGetUsers struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Meta    []string `json:"meta"`
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users in DB.
// @Produce application/json
// @Tags account/coins
// @Success 200 {object} ResponseGetUsers{}
// @Router /v1/account/coins [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var database *tools.DatabaseInterface = tools.NewDatabase()
	var users []string = (*database).GetUsers()

	api.SuccessResponse(w, users, "Users Accepted", http.StatusOK, "success")
}

type ResponseGetBalance struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Meta    api.CoinBalanceResponse `json:"meta"`
}

// GetCoinBalance godoc
// @Summary Get user balance
// @Description Get user balance in DB.
// @Param username path string true "Username"
// @Produce application/json
// @Tags account/coins
// @Success 200 {object} ResponseGetBalance{}
// @Router /v1/account/coins/{username} [get]
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

type ResponseUpdateCoinBalance struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Meta    tools.CoinDetails `json:"meta"`
}

// UpdateCoinBalance godoc
// @Summary Update user balance
// @Description Update user balance in DB.
// @Param username path string true "Username"
// @Param body body api.UpdateCoinBalanceBody true "Body"
// @Produce application/json
// @Tags account/coins
// @Success 200 {object} ResponseUpdateCoinBalance{}
// @Router /v1/account/coins/{username} [put]
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

type ResponseNewUser struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Meta    api.NewUserBody `json:"meta"`
}

// NewUser godoc
// @Summary Create new user
// @Description Create new user in DB.
// @Param body body api.NewUserBody true "Body"
// @Produce application/json
// @Tags account/coins
// @Success 201 {object} ResponseNewUser{}
// @Router /v1/account/coins/{username} [post]
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

type ResponseRemoveUser struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Meta    api.CoinBalanceParams `json:"meta"`
}

// RemoveUser godoc
// @Summary Remove user
// @Description Remove user in DB.
// @Param username path string true "Username"
// @Produce application/json
// @Tags account/coins
// @Success 202 {object} RemoveUser{}
// @Router /v1/account/coins/{username} [delete]
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
