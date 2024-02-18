package routes

import (
	"Raghava/OneCV-Assignment/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func GetAllRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/register", handlers.Register)
		r.Get("/commonstudents", handlers.CommonStudents)
		r.Post("/suspend", handlers.Suspend)
		r.Post("/retrievefornotifications", handlers.RetrieveForNotifications)
	}
}
