package server

import (
	"fmt"
	"net/http"
	"pajarit-feed-service/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func StartServer(cfg *config.Configuration, deps *config.Dependencies) error {

	// Quizá para el challenge es un poco excesivo, pero el uso de go-chi me parece una buena
	// opción ya que es un poco más declarativo que la implementación vanilla, permite un manejo
	// más fácil de los params y es liviana.
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/posts", CreatePostHandler(deps))
		r.Post("/followup", FollowUserHandler(deps))
		r.Get("/timelines/{userId}", GetTimelineHandler(deps))
	})

	serverPort := fmt.Sprintf(":%d", cfg.ServerPort)
	return http.ListenAndServe(serverPort, r)
}
