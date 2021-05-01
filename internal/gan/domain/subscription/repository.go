package subscription

import (
	"github.com/v4rakh/gan/internal/gan/domain"
	"gorm.io/gorm"
)

type repository interface {
	Paginate(page int, pageSize int, orderBy string, order string) ([]*Subscription, error)
	ListWhereState(state State) ([]*Subscription, error)
	Count() (int64, error)
	Find(address string) (*Subscription, error)
	Create(address string, state State, token string) error
	Update(address string, state State, token string) (*Subscription, error)
	Delete(address string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) Find(address string) (*Subscription, error) {
	var sub Subscription

	res := r.db.Find(&sub, "address = ?", address)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, domain.ErrorNotFound
	}

	return &sub, nil
}

func (r *repo) Create(address string, state State, token string) error {
	if address == "" || token == "" {
		return domain.ErrorValidationNotBlank
	}

	var e *Subscription

	e = &Subscription{
		Address: address,
		State:   state.Value(),
		Token:   token,
	}

	res := r.db.Create(&e)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return domain.ErrorCreateFailed
	}

	return nil
}

func (r *repo) Update(address string, state State, token string) (*Subscription, error) {
	if address == "" || token == "" {
		return nil, domain.ErrorValidationNotBlank
	}

	e, err := r.Find(address)

	if err != nil {
		return nil, err
	}

	e.State = state.Value()
	e.Token = token

	res := r.db.Save(&e)

	if res.RowsAffected == 0 {
		return e, domain.ErrorUpdateFailed
	}

	return e, nil
}

func (r *repo) Delete(address string) error {
	if address == "" {
		return domain.ErrorValidationNotBlank
	}

	res := r.db.Delete(&Subscription{}, "address = ?", address)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return domain.ErrorDeleteFailed
	}

	return nil
}

func (r *repo) Paginate(page int, pageSize int, orderBy string, order string) ([]*Subscription, error) {
	var e []*Subscription

	if page == 0 || pageSize <= 0 {
		return nil, domain.ErrorValidationPageGreaterZero
	}

	if pageSize <= 0 {
		return nil, domain.ErrorValidationPageSizeGreaterZero
	}

	offset := (page - 1) * pageSize

	var res *gorm.DB
	if orderBy != "" && order != "" {
		res = r.db.Order(orderBy + " " + order).Offset(offset).Limit(pageSize).Find(&e)
	} else {
		res = r.db.Offset(offset).Limit(pageSize).Find(&e)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return e, nil
}

func (r *repo) ListWhereState(state State) ([]*Subscription, error) {
	var e []*Subscription
	var res *gorm.DB

	res = r.db.Where("state = ?", state.Value()).Find(&e)

	if res.Error != nil {
		return nil, res.Error
	}

	return e, nil
}

func (r *repo) Count() (int64, error) {
	var c int64

	var res *gorm.DB
	res = r.db.Model(&Subscription{}).Count(&c)

	if res.Error != nil {
		return 0, res.Error
	}

	return c, nil
}
