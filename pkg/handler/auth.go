package handler

import (
	"github.com/TimurMamdov1991/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) singUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"message": "Добро пожаловать! Это сервис LifePainQA",
	})
}

func (h *Handler) singIn(c *gin.Context) {

}
