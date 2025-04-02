package server

import (
	"fmt"
	"net/http"
	"pajarit-feed-service/config"
)

func StartServer(cfg *config.Configuration, deps *config.Dependencies) error {

	// Opté por levantar un server vanilla; quizá chi hubiese sido un poco más prolijo
	// pero son pocas rutas las que tengo que disponibilizar y, además, me tomo la
	// libertad de prescindir de usar middlewares

	appServer := http.NewServeMux()
	appServer.HandleFunc("api/v1/posts", CreatePostHandler())
	appServer.HandleFunc("api/v1/followup", FollowUserHandler())
	appServer.HandleFunc("api/v1/timelines", GetTimelineHandler())

	serverPort := fmt.Sprintf(":%d", cfg.Port)
	return http.ListenAndServe(serverPort, appServer)
}
