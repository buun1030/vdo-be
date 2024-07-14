package video

type Video struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	UserID int64  `json:"userID"`
}
