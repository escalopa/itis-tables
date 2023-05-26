package core

type Subject struct {
	From  string      `json:"from"`
	To    string      `json:"to"`
	Title string      `json:"title"`
	Meta  SubjectMeta `json:"meta"`
}

type SubjectMeta struct {
	RoomNumber string `json:"room_number"`
	Location   string `json:"location"`
	Teacher    string `json:"teacher"`
}
