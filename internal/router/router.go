package router

import (
	"Raghava/OneCV-Assignment/internal/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Setup() chi.Router {
	r := chi.NewRouter()
	setUpRoutes(r)
	return r
}

// Set up middleware and routes
func setUpRoutes(r chi.Router) {
	r.Use(middleware.Recoverer)
	r.Route("/api", func(r chi.Router) {
		r.Group(routes.GetAllRoutes())
	})
}
