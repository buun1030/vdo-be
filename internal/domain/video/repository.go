package video

type VideoRepository interface {
	CreateVideo(v *Video) error
}
