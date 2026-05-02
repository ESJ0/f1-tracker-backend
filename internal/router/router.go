package router

import (
	"database/sql"
	"net/http"

	"f1-tracker-backend/internal/handlers"
	"f1-tracker-backend/internal/middleware"
	"f1-tracker-backend/internal/repository"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func New(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	// Middlewares globales
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.CORS)

	// Repos y handlers
	driverRepo := repository.NewDriverRepo(db)
	driverHandler := handlers.NewDriverHandler(driverRepo)

	// Rutas
	r.Route("/drivers", func(r chi.Router) {
		r.Get("/", driverHandler.GetAll)
		r.Post("/", driverHandler.Create)
		r.Get("/{id}", driverHandler.GetByID)
		r.Put("/{id}", driverHandler.Update)
		r.Delete("/{id}", driverHandler.Delete)
	})

	return r
}
