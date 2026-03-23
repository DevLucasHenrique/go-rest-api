package handlers

import (
	"github.com/go-chi"
	chimiddle "github.com/go-chil/middleware"
	"github.com/avukadin/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global Middleware
	r.Use(chimiddle.StripSlashes)

	r.Use("/account", func(router, chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}