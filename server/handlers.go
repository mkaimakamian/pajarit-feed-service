package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	createpost "pajarit-feed-service/application/create_post"
	followuser "pajarit-feed-service/application/follow_user"
	gettimeline "pajarit-feed-service/application/get_timeline"
	"pajarit-feed-service/config"
	"strconv"

	"github.com/go-chi/chi"
)

func CreatePostHandler(deps *config.Dependencies) http.HandlerFunc {
	usecase := createpost.NewCreatePost(deps.PostRepository, deps.EventPublisher)

	return func(w http.ResponseWriter, r *http.Request) {
		var cmd createpost.CreatePostCmd
		if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
			HttpBadRequestError(w, err)
			return
		}

		response, err := usecase.Exec(r.Context(), cmd)
		if err != nil {
			HttpInternalServerError(w, err)
			return
		}

		HttpCreated(w, response)
	}
}

func FollowUserHandler(deps *config.Dependencies) http.HandlerFunc {
	usecase := followuser.NewFollowUser(deps.FollowUpRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		var cmd followuser.FollowUsertCmd
		if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
			HttpBadRequestError(w, err)
			return
		}

		if err := usecase.Exec(r.Context(), cmd); err != nil {
			HttpInternalServerError(w, err)
			return
		}

		HttpCreated(w, nil)
	}
}

func GetTimelineHandler(deps *config.Dependencies) http.HandlerFunc {
	usecase := gettimeline.NewGetTimeline(deps.TimelineRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		if userId == "" {
			HttpBadRequestError(w, fmt.Errorf("userId is required"))
			return
		}

		offsetParam := r.URL.Query().Get("offset")
		sizeParam := r.URL.Query().Get("size")

		offset, _ := strconv.Atoi(offsetParam)
		size, _ := strconv.Atoi(sizeParam)

		cmd := gettimeline.GetTimelineCmd{UserId: userId, Offset: offset, Size: size}

		response, err := usecase.Exec(r.Context(), cmd)
		if err != nil {
			HttpInternalServerError(w, err)
			return
		}

		HttpOk(w, response)
	}
}
