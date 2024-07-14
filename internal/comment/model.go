package comment

type Comment struct {
	ID      int64  `json:"id"`
	VideoID int64  `json:"videoID"`
	UserID  int64  `json:"userID"`
	Content string `json:"content"`
}
