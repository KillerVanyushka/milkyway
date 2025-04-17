package models

type PostEdit struct {
	Content string `json:"content"`
	UserId  int    `json:"userId"`
}
