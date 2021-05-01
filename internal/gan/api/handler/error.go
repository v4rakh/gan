package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/v4rakh/gan/internal/gan/api/presenter"
	"github.com/v4rakh/gan/internal/gan/domain"
	"net/http"
)

func HandleAbortWhenError(c *gin.Context, err error) {
	if err != nil {
		if err == domain.ErrorValidationNotBlank || err == domain.ErrorValidationPageGreaterZero || err == domain.ErrorValidationPageSizeGreaterZero {
			c.AbortWithStatusJSON(http.StatusBadRequest, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorBadRequest, err.Error()))
			return
		} else if err == domain.ErrorForbiddenTokenMatch {
			c.AbortWithStatusJSON(http.StatusForbidden, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorForbidden, err.Error()))
			return
		} else if err == domain.ErrorNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorNotFound, err.Error()))
			return
		} else if err == domain.ErrorConflict {
			c.AbortWithStatusJSON(http.StatusConflict, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorConflict, err.Error()))
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.NewErrorResponseWithMessage(err.Error()))
			return
		}
	}
}
