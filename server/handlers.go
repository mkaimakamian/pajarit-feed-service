package server

import (
	"encoding/json"
	"net/http"
	createpost "pajarit-feed-service/application/create_post"
	followuser "pajarit-feed-service/application/follow_user"
	gettimeline "pajarit-feed-service/application/get_timeline"
	"pajarit-feed-service/config"
)

func CreatePostHandler(deps *config.Dependencies) http.HandlerFunc {

	usecase := createpost.NewCreatePost(deps.PostRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		// Con chi u otro enfoque que filtre por tipo de método, no sería necesario
		// este tipo de exclusión.
		if r.Method != http.MethodPost {
			HttpMethodNotAllowed(w)
		}

		var cmd createpost.CreatePostCmd
		err := json.NewDecoder(r.Body).Decode(&cmd)
		if err != nil {
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
		if r.Method != http.MethodPost {
			HttpMethodNotAllowed(w)
		}

		var cmd followuser.FollowUsertCmd
		err := json.NewDecoder(r.Body).Decode(&cmd)
		if err != nil {
			HttpBadRequestError(w, err)
			return
		}

		err = usecase.Exec(r.Context(), cmd)
		if err != nil {
			HttpInternalServerError(w, err)
			return
		}

		HttpCreated(w, "")
	}

}

func GetTimelineHandler(deps *config.Dependencies) http.HandlerFunc {
	usecase := gettimeline.NewGetTimeline(deps.TimelineRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		// Con chi u otro enfoque que filtre por tipo de método, no sería necesario
		// este tipo de exclusión.
		if r.Method != http.MethodPost {
			HttpMethodNotAllowed(w)
		}

		var cmd gettimeline.GetTimelineCmd
		err := json.NewDecoder(r.Body).Decode(&cmd)
		if err != nil {
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
