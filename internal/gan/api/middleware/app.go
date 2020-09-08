package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/v4rakh/gan/internal/gan/constant"
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
