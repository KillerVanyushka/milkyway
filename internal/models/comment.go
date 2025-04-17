package models

type Comment struct {
	Id     int    `json:"id"`
	PostId int    `json:"postId"`
	UserId int    `json:"userId"`
	Text   string `json:"text"`
}
