package handler

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"reflect"
// 	"strings"
// 	"vdo-be/pkg/api"
// )

// const (
// 	mimeApplicationJSON   = "application/json"
// 	mimeMultipartFormData = "multipart/form-data"
// )

// func decode[T any](r *http.Request) (T, error) {
// 	var v T
// 	contentType := r.Header.Get("Content-Type")
// 	switch {
// 	case contentType == mimeApplicationJSON:
// 		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
// 			return v, fmt.Errorf("decode json: %w", err)
// 		}
// 	case strings.HasPrefix(contentType, mimeMultipartFormData):
// 		f, h, err := r.FormFile("video")
// 		if err != nil {
// 			return v, fmt.Errorf("get form file: %w", err)
// 		}
// 		file := api.File{
// 			File:       f,
// 			FileHeader: h,
// 		}

// 		v = reflect.ValueOf(file)

// 	}
// 	return v, nil
// }
