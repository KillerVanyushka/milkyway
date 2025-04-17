package delivery

import (
	"github.com/gin-gonic/gin"
	"milkyway/internal/models"
	"milkyway/internal/services"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service *services.UserService
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
func (h *UserHandler) GetUserByid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неверный айди пользователя"})
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Ошибка": "Пользователь не найден"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userCreate models.UserEdit
	if err := c.BindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неверное тело запроса"})
	}
	newUser := h.service.CreateUser(userCreate)
	c.JSON(http.StatusOK, newUser)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неверный айди пользователя"})
		return
	}

	var userEdit models.UserEdit
	if err := c.ShouldBindJSON(&userEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неверное тело запроса"})
		return
	}

	updatedUser, err := h.service.UpdateUser(id, userEdit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Ошибка": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)

}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Неверный айди пользователя"})
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Ошибка": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Уведомление": "Пользователь удален"})
}
