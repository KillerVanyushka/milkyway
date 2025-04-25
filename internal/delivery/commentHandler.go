package delivery

import (
	"github.com/gin-gonic/gin"
	"milkyway/internal/models"
	"milkyway/internal/services"
	"net/http"
	"strconv"
)

func NewCommentHandler(service *services.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

type CommentHandler struct {
	service *services.CommentService
}

func (h *CommentHandler) GetAllComments(c *gin.Context) {
	comments, _ := h.service.GetAllComments()
	c.JSON(http.StatusOK, comments)
}

func (h *PostHandler) GetCommentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неверный айди поста"})
		return
	}

	comment, err := h.service.GetPostById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Ошибка": "Пост не найден"})
		return
	}

	c.JSON(http.StatusOK, comment)

}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var commentCreate models.CommentEdit
	if err := c.ShouldBindJSON(&commentCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неверное тело запроса"})
	}

	newComment, err := h.service.CreateComment(commentCreate.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Ошибка": "Не получилось добавить"})
	}
	c.JSON(http.StatusOK, newComment)
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неверный айди комментарий"})
		return
	}

	var commentEdit models.CommentEdit
	if err := c.ShouldBindJSON(&commentEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неверное тело запроса"})
		return
	}

	updatedComment, err := h.service.UpdateComment(id, &commentEdit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Ошибка": "Комментарий не найден"})
		return
	}

	c.JSON(http.StatusOK, updatedComment)
}
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неверный айди комментарий"})
		return
	}

	if err := h.service.DeleteComment(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Ошибка": "Комментарий не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Уведомление": "Комментарий удален"})
}
