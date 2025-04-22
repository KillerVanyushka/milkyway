package repository

import (
	"gorm.io/gorm"
	"milkyway/internal/models"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{db: db}
}

func (s UserRepoImpl) GetAll() ([]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error
	return users, err
}
func (s UserRepoImpl) GetById(id int) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	return &user, err
}

func (s UserRepoImpl) Create(user models.User) error {
	return s.db.Create(&user).Error
}

func (s UserRepoImpl) Update(id int, user *models.UserEdit) error {
	return s.db.Model(&models.User{}).Where("id=?", id).Omit("id, CreateAt").Updates(user).Error
}

func (s UserRepoImpl) Delete(id int) error {
	return s.db.Delete(&models.User{}, id).Error
}
