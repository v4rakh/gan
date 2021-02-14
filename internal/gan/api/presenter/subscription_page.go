package presenter

type SubscriptionPage struct {
	Content       []*Subscription `json:"content"`
	Page          int             `json:"page"`
	PageSize      int             `json:"pageSize"`
	OrderBy       string          `json:"orderBy"`
	Order         string          `json:"order"`
	TotalElements int64           `json:"totalElements"`
	TotalPages    int64           `json:"totalPages"`
}

func NewSubscriptionPage(content []*Subscription, page int, pageSize int, orderBy string, order string, totalElements int64, totalPages int64) *SubscriptionPage {
	e := new(SubscriptionPage)
	e.Content = content
	e.Page = page
	e.PageSize = pageSize
	e.OrderBy = orderBy
	e.Order = order
	e.TotalElements = totalElements
	e.TotalPages = totalPages
	return e
}
