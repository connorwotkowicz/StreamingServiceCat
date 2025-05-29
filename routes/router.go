// This defines all API routes and returns full router

package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"streamflix-api/handlers"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter() 
	//  'chi' is a lightweight router for Go


	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)


	// routes
	r.Get("/media", handlers.GetAllMedia)
	r.Get("/media/{id}", handlers.GetMediaByID)
	r.Post("/media", handlers.CreateMedia)
	r.Put("/media/{id}", handlers.UpdateMedia)
	r.Delete("/media/{id}", handlers.DeleteMedia)
	// Maps endpoints like /media to handler functions in the handlers package (which we’ll build next).

	return r
}
