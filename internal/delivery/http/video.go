package handler

import (
	"context"
	"net/http"
	"vdo-be/internal/domain/video"
	"vdo-be/internal/infra/storage/gcp"
)

func uploadVideo(w http.ResponseWriter, r *http.Request) {
	// Parse form data and extract video file
	file, _, err := r.FormFile("video")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create video domain object with metadata
	video := &video.Video{
		Title:  r.FormValue("title"),
		UserID: 1,
		Metadata: map[string]string{
			"category": r.FormValue("category"),
			"tags":     r.FormValue("tags"), // Might need further processing
		},
	}

	ctx := context.Background()

	// Upload video to cloud storage
	videoID, err := gcp.UploadVideo(ctx, file, video.Metadata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	video.UpdateID(videoID)
}
