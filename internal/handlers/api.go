package handlers

import (
	"github.com/DevLucasHenrique/go-rest-api/internal/handlers/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global Middleware
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
