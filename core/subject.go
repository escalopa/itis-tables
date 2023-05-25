package core

type Subject struct {
	From     string   `json:"from"`
	To       string   `json:"to"`
	Course   Course   `json:"course"`
	WeekType WeekType `json:"week_type"`
}
