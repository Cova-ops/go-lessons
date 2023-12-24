package handlers

import (
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"

	"go-lessons/api"
	"go-lessons/internal/handlers/account"
	"go-lessons/internal/middlewares"
)

const URL_PREFIX = "/api/v1"

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(api.ErrorHandler)
	r.Use(chimiddle.StripSlashes)

	r.Route(URL_PREFIX+"/account", func(r chi.Router) {

		r.Get("/coins", account.GetUsers)

		r.With(middlewares.AuthorizationMiddleware).Get("/coins/{username}", account.GetCoinBalance)
		r.With(middlewares.AuthorizationMiddleware).Put("/coins/{username}", account.UpdateCoinBalance)
		r.With(middlewares.AuthorizationMiddleware).Post("/coins/{username}", account.NewUser) // Only a existing user can create a new use
		r.With(middlewares.AuthorizationMiddleware).Delete("/coins/{username}", account.RemoveUser)
	})
}
