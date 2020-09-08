package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/v4rakh/gan/internal/gan/api"
	"github.com/v4rakh/gan/internal/gan/api/presenter"
	"github.com/v4rakh/gan/internal/gan/domain"
	"github.com/v4rakh/gan/internal/gan/domain/announcement"
	"github.com/v4rakh/gan/internal/util"
	"net/http"
	"strings"
)

type AnnouncementHandler struct {
	service announcement.Service
}

func NewAnnouncementHandler(s *announcement.Service) *AnnouncementHandler {
	return &AnnouncementHandler{service: *s}
}

type createAnnouncementRequest struct {
	Title   string `json:"title" binding:"required,min=3,max=255"`
	Content string `json:"content" binding:"required,min=1"`
}

type updateAnnouncementRequest struct {
	ID      string `json:"id" binding:"required,min=1"`
	Title   string `json:"title" binding:"required,min=3,max=255"`
	Content string `json:"content" binding:"required,min=1"`
}

func (h *AnnouncementHandler) ListAnnouncements(c *gin.Context) {
	orderBy := util.ToSnakeCase(c.Query("sortBy"))
	order := strings.ToLower(c.Query("sortOrder"))

	if orderBy == "" || !util.FindInSlice([]string{"id", "created_at"}, orderBy) {
		orderBy = "created_at"
	}

	if order == "" {
		order = "desc"
	}

	announcements, err := h.service.List(orderBy, order)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.NewErrorResponseWithMessage(err.Error()))
		return
	}

	var data []*presenter.Announcement
	data = make([]*presenter.Announcement, 0)

	for _, e := range announcements {
		data = append(data, &presenter.Announcement{
			ID:        e.ID,
			Title:     e.Title,
			Content:   e.Content,
			CreatedAt: e.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, presenter.NewDataResponseWithPayload(data))
}

func (h *AnnouncementHandler) GetAnnouncement(c *gin.Context) {
	e, err := h.service.Get(c.Param("id"))

	if err != nil {
		if err == domain.ErrorAnnouncementNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorNotFound, err.Error()))
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.NewErrorResponseWithMessage(err.Error()))
			return
		}
	}

	data := &presenter.Announcement{
		ID:        e.ID,
		Title:     e.Title,
		Content:   e.Content,
		CreatedAt: e.CreatedAt,
	}

	c.JSON(http.StatusOK, presenter.NewDataResponseWithPayload(data))
}

func (h *AnnouncementHandler) CreateAnnouncement(c *gin.Context) {
	var req createAnnouncementRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		res := api.ConvertErrorsTo(&errs)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	e, err := h.service.Create(req.Title, req.Content)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.NewErrorResponseWithMessage(err.Error()))
		return
	}

	data := &presenter.Announcement{
		ID:        e.ID,
		Title:     e.Title,
		Content:   e.Content,
		CreatedAt: e.CreatedAt,
	}

	c.JSON(http.StatusOK, presenter.NewDataResponseWithPayload(data))
}

func (h *AnnouncementHandler) UpdateAnnouncement(c *gin.Context) {
	var req updateAnnouncementRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		res := api.ConvertErrorsTo(&errs)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	e, err := h.service.Update(req.ID, req.Title, req.Content)

	if err != nil {
		if err == domain.ErrorAnnouncementNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorNotFound, err.Error()))
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.NewErrorResponseWithMessage(err.Error()))
			return
		}
	}

	data := &presenter.Announcement{
		ID:        e.ID,
		Title:     e.Title,
		Content:   e.Content,
		CreatedAt: e.CreatedAt,
	}

	c.JSON(http.StatusOK, presenter.NewDataResponseWithPayload(data))
}

func (h *AnnouncementHandler) DeleteAnnouncement(c *gin.Context) {
	err := h.service.Delete(c.Param("id"))

	if err != nil {
		if err == domain.ErrorAnnouncementNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, presenter.NewErrorResponseWithStatusAndMessage(presenter.ErrorNotFound, err.Error()))
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.NewErrorResponseWithMessage(err.Error()))
			return
		}
	}

	c.JSON(http.StatusNoContent, nil)
}
