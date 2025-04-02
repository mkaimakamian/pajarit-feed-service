package server

import (
	"fmt"
	"net/http"
	"pajarit-feed-service/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func StartServer(cfg *config.Configuration, deps *config.Dependencies) error {

	// Uso chi para poder manejar más fácilmente las rutas y poder obtener los pathparams
	// Quizá para el challenge es un poco excesivo, pero es un poco más declarativo y está más ordenadoi.
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/posts", CreatePostHandler(deps))
		r.Post("/followup", FollowUserHandler(deps))
		r.Get("/timelines/{userId}", GetTimelineHandler(deps))
	})

	serverPort := fmt.Sprintf(":%d", cfg.Port)
	return http.ListenAndServe(serverPort, r)
}
