package handler

import (
	"encoding/json"
	"net/http"
)

func uploadVideo() http.Handler {
	type response struct {
		Greeting string `json:"greeting"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse form data and extract video file
		// file, _, err := r.FormFile("video")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// defer file.Close()

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

		if err := json.NewEncoder(w).Encode(response{Greeting: "Upload video"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
