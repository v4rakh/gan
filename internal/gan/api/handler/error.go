package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/v4rakh/gan/internal/errors"
	"github.com/v4rakh/gan/internal/gan/api/presenter"
	"net/http"
)

func HandleAbortWhenError(c *gin.Context, err error) {
	if c == nil || err == nil {
		return
	}

	switch e := err.(type) {
	case *errors.ServiceError:
		if e.Status == errors.IllegalArgument {
			c.AbortWithStatusJSON(http.StatusBadRequest, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorIllegalArgument, e.Error()))
			return
		} else if e.Status == errors.Unauthorized {
			c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorUnauthorized, e.Error()))
			return
		} else if e.Status == errors.Forbidden {
			c.AbortWithStatusJSON(http.StatusForbidden, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorForbidden, e.Error()))
			return
		} else if e.Status == errors.NotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorNotFound, e.Error()))
			return
		} else if e.Status == errors.Conflict {
			c.AbortWithStatusJSON(http.StatusConflict, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorConflict, e.Error()))
			return
		} else if e.Status == errors.GeneralError {
			c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.NewErrorResponseWithMessage(e.Error()))
			return
		}
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.NewErrorResponseWithMessage(err.Error()))
		return
	}
}
