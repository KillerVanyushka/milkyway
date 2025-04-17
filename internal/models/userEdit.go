package models

type UserEdit struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
}
