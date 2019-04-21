package routes

import (
	"github.com/go-chi/chi"
	"redis_proxy/app/handlers"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{postID}", handlers.HandleGet)
	router.Post("/", handlers.HandleCreate)
	return router
}
