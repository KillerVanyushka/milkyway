package models

type CommentEdit struct {
	Text   string `json:"text"`
	PostId int    `json:"postId"`
	UserId int    `json:"userId"`
}
