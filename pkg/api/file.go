package api

import "mime/multipart"

type File struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}
