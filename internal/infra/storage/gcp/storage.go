package gcp

import (
	"context"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

// UploadVideo uploads a video file to Google Cloud Storage
func UploadVideo(ctx context.Context, file multipart.File, metadata map[string]string) (string, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}

	// Upload video file to Cloud Storage
	wc := client.Bucket("video-bucket").Object(uuid.New().String()).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	return wc.Attrs().Name, nil
}
