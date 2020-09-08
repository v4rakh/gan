package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/v4rakh/gan/internal/gan/constant"
	"net/http"
)

type InfoHandler struct {
}

func NewInfoHandler() *InfoHandler {
	return &InfoHandler{}
}

func (h *InfoHandler) ShowInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    constant.AppName,
		"version": constant.AppVersion,
	})
}
