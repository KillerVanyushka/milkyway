package models

type UserMig struct {
	Id          uint   `gorm:"primaryKey"`
	Username    string `gorm:"unique;not null"`
	PhoneNumber string `gorm:"not null"`
	Age         int    `gorm:""`
	Gender      string `gorm:""`
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
}
