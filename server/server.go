package server

import (
	"fmt"
	"net/http"
	"pajarit-feed-service/config"
)

func StartServer(cfg *config.Configuration, deps *config.Dependencies) error {

	// Opté por levantar un server vanilla; quizá chi hubiese sido un poco más prolijo
	// pero son pocas rutas las que tengo que disponibilizar, se emplea un único método por recurso
	// y, además, presciendo del uso de middlewares

	appServer := http.NewServeMux()
	appServer.HandleFunc("api/v1/posts", CreatePostHandler(deps))
	appServer.HandleFunc("api/v1/followup", FollowUserHandler(deps))
	appServer.HandleFunc("api/v1/timelines", GetTimelineHandler(deps))

	serverPort := fmt.Sprintf(":%d", cfg.Port)
	return http.ListenAndServe(serverPort, appServer)
}
