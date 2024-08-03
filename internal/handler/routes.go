package handler

import (
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
) {
	mux.Handle("POST /upload", uploadVideo())
}
