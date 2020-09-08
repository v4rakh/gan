package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Login(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
