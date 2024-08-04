package handler

import (
	"fmt"
	"net/http"
	"vdo-be/pkg/api"
)

func uploadVideo() http.Handler {

	type request struct {
		FileName string `json:"fileName"`
		Title    string `json:"title"`
		Category string `json:"category"`
		Tags     string `json:"tags"`
	}

	type response struct {
		FileName string `json:"fileName"`
		Title    string `json:"title"`
		Category string `json:"category"`
		Tags     string `json:"tags"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request
		req, err := decode[request](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create video domain object with metadata
		// video := &video.Video{
		// 	Title:  r.FormValue("title"),
		// 	UserID: 1,
		// 	Metadata: map[string]string{
		// 		"category": r.FormValue("category"),
		// 		"tags":     r.FormValue("tags"), // Might need further processing
		// 	},
		// }

		// ctx := context.Background()

		// Upload video to cloud storage
		// videoID, err := gcp.UploadVideo(ctx, file, video.Metadata)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// video.UpdateID(videoID)

		// insert video to database
		// if err := command.CreateVideo(video); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		resp := response{
			FileName: req.FileName,
			Title:    req.Title,
			Category: req.Category,
			Tags:     fmt.Sprintf("Tags: %s", req.Tags),
		}

		if err := api.WriteResponse(w, http.StatusCreated, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
