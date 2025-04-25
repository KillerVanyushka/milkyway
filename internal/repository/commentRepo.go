package repository

import (
	"gorm.io/gorm"
	"milkyway/internal/models"
)

type CommentRepoImpl struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepoImpl {
	return &CommentRepoImpl{db: db}
}

func (c CommentRepoImpl) GetAll() ([]models.Comment, error) {
	var comments []models.Comment
	err := c.db.Find(&comments).Error
	return comments, err
}
func (c CommentRepoImpl) GetById(id int) (*models.Comment, error) {
	var comment models.Comment
	err := c.db.First(&comment, id).Error
	return &comment, err
}

func (c CommentRepoImpl) Create(comment models.Comment) error {
	return c.db.Create(&comment).Error
}

func (c CommentRepoImpl) Update(id int, comment *models.CommentEdit) error {
	return c.db.Model(&models.Comment{}).Where("id=?", id).Omit("id, CreateAt").Updates(comment).Error
}

func (c CommentRepoImpl) Delete(id int) error {
	return c.db.Delete(&models.Comment{}, id).Error
}
