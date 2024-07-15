package video

type Video struct {
	ID       string            `json:"id"`
	Title    string            `json:"title"`
	UserID   int64             `json:"userID"`
	Metadata map[string]string `json:"metadata"`
}

func (v *Video) UpdateID(id string) {
	v.ID = id
}
