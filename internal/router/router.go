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

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.CORS)

	// Repos
	driverRepo := repository.NewDriverRepo(db)
	raceRepo := repository.NewRaceRepo(db)
	resultRepo := repository.NewResultRepo(db)

	// Handlers
	driverHandler := handlers.NewDriverHandler(driverRepo)
	raceHandler := handlers.NewRaceHandler(raceRepo)
	resultHandler := handlers.NewResultHandler(resultRepo)

	// Driver routes
	r.Route("/drivers", func(r chi.Router) {
		r.Get("/", driverHandler.GetAll)
		r.Post("/", driverHandler.Create)
		r.Get("/{id}", driverHandler.GetByID)
		r.Put("/{id}", driverHandler.Update)
		r.Delete("/{id}", driverHandler.Delete)
		r.Get("/{id}/results", resultHandler.GetByDriver) // resultados de un piloto
	})

	// Race routes
	r.Route("/races", func(r chi.Router) {
		r.Get("/", raceHandler.GetAll)
		r.Post("/", raceHandler.Create)
		r.Get("/{id}", raceHandler.GetByID)
		r.Put("/{id}", raceHandler.Update)
		r.Delete("/{id}", raceHandler.Delete)
		r.Get("/{id}/results", resultHandler.GetByRace) // resultados de una carrera
	})

	// Results routes
	r.Route("/results", func(r chi.Router) {
		r.Post("/", resultHandler.Create)
		r.Delete("/{id}", resultHandler.Delete)
	})

	return r
}
