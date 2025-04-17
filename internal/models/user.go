package models

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
