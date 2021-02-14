package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/v4rakh/gan/internal/gan/api"
	"github.com/v4rakh/gan/internal/gan/api/presenter"
	"github.com/v4rakh/gan/internal/gan/domain/subscription"
	"github.com/v4rakh/gan/internal/util"
	"net/http"
	"strconv"
	"strings"
)

type SubscriptionHandler struct {
	service subscription.Service
}

func NewSubscriptionHandler(s *subscription.Service) *SubscriptionHandler {
	return &SubscriptionHandler{service: *s}
}

type createSubscriptionRequest struct {
	Address string `json:"address" binding:"required,min=3,max=255"`
}

type rescueSubscriptionRequest struct {
	Address string `json:"address" binding:"required,min=3,max=255"`
}

type verifySubscriptionRequest struct {
	Address string `json:"address" binding:"required,min=1"`
	Token   string `json:"token" binding:"required,min=1"`
}

type deleteSubscriptionRequest struct {
	Address string `json:"address" binding:"required,min=1"`
	Token   string `json:"token" binding:"required,min=1"`
}

type deleteSubscriptionByAddressRequest struct {
	Address string `json:"address" binding:"required,min=1"`
}

func (h *SubscriptionHandler) PaginateSubscriptions(c *gin.Context) {
	orderBy := util.ToSnakeCase(c.Query("orderBy"))
	order := strings.ToLower(c.Query("order"))
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	if page == 0 {
		page = 0
	}

	if pageSize <= 0 {
		pageSize = 5
	}

	if orderBy == "" || !util.FindInSlice([]string{"address", "created_at", "updated_at"}, orderBy) {
		orderBy = "created_at"
	}

	if order == "" {
		order = "desc"
	}

	subscriptions, err := h.service.Paginate(page+1, pageSize, orderBy, order)
	if err != nil {
		HandleAbortWhenError(c, err)
		return
	}

	var data []*presenter.Subscription
	data = make([]*presenter.Subscription, 0)

	for _, e := range subscriptions {
		data = append(data, &presenter.Subscription{
			Address:   e.Address,
			Token:     e.Token,
			State:     e.State,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		})
	}

	totalElements, err := h.service.Count()
	if err != nil {
		HandleAbortWhenError(c, err)
		return
	}

	totalPages := (totalElements + int64(pageSize) - 1) / int64(pageSize)
	c.JSON(http.StatusOK, presenter.NewDataResponseWithPayload(presenter.NewSubscriptionPage(data, page, pageSize, orderBy, order, totalElements, totalPages)))
}

func (h *SubscriptionHandler) CreateSubscription(c *gin.Context) {
	var req createSubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		res := api.ConvertErrorsTo(&errs)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := h.service.Create(req.Address)
	if err != nil {
		HandleAbortWhenError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *SubscriptionHandler) VerifySubscription(c *gin.Context) {
	var req verifySubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		res := api.ConvertErrorsTo(&errs)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := h.service.Verify(req.Address, req.Token)
	if err != nil {
		HandleAbortWhenError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *SubscriptionHandler) RescueSubscription(c *gin.Context) {
	var req rescueSubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		res := api.ConvertErrorsTo(&errs)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := h.service.Rescue(req.Address)
	if err != nil {
		HandleAbortWhenError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *SubscriptionHandler) DeleteSubscription(c *gin.Context) {
	var req deleteSubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		res := api.ConvertErrorsTo(&errs)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := h.service.Delete(req.Address, req.Token)
	if err != nil {
		HandleAbortWhenError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *SubscriptionHandler) DeleteSubscriptionByAddress(c *gin.Context) {
	var req deleteSubscriptionByAddressRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		res := api.ConvertErrorsTo(&errs)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := h.service.DeleteByAddress(req.Address)
	if err != nil {
		HandleAbortWhenError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
