package database

import "vdo-be/internal/domain/video"

type Command interface {
	CreateVideo(v *video.Video) error
}
