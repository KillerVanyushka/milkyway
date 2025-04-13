package models

type PostEdit struct {
	Content string `json:"content"`
	UserId  string `json:"userId"`
}
