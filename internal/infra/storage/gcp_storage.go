package storage

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/google/uuid"
)

// UploadVideo uploads a video file to Google Cloud Storage
func UploadVideo(ctx context.Context, file multipart.File, metadata map[string]string) (string, error) {
	client, err := NewClient(ctx)
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
