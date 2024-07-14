package handler

import (
	"context"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func uploadVideo(w http.ResponseWriter, r *http.Request) {
	// Parse form data and extract video file
	file, _, err := r.FormFile("video")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a Cloud Storage client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Upload video file to Cloud Storage
	wc := client.Bucket("video-bucket").Object(uuid.New().String()).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := wc.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
