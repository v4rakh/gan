package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/v4rakh/gan/internal/gan/api/presenter"
	"github.com/v4rakh/gan/internal/gan/constant"
	"net/http"
)

func AppName() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(constant.AppNameHeader, constant.AppName)
		c.Next()
	}
}

func AppVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(constant.AppVersionHeader, constant.AppVersion)
		c.Next()
	}
}

func AppErrorRecoveryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			c.Header(constant.AppVersionHeader, constant.AppVersion)
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.NewErrorResponseWithMessage(fmt.Sprintf("%s", err)))
			}
		}()
		c.Next()
	}
}
