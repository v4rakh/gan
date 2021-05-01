package announcement

import (
	"github.com/v4rakh/gan/internal/gan/domain"
	"gorm.io/gorm"
)

type repository interface {
	Paginate(page int, pageSize int, orderBy string, order string) ([]*Announcement, error)
	Count() (int64, error)
	Find(id string) (*Announcement, error)
	Create(title string, content string) (*Announcement, error)
	Update(id string, title string, content string) (*Announcement, error)
	Delete(id string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) Find(id string) (*Announcement, error) {
	if id == "" {
		return nil, domain.ErrorValidationNotBlank
	}

	var announcement Announcement

	res := r.db.Find(&announcement, "id = ?", id)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, domain.ErrorNotFound
	}

	return &announcement, nil
}

func (r *repo) Create(title string, content string) (*Announcement, error) {
	if title == "" || content == "" {
		return nil, domain.ErrorValidationNotBlank
	}

	var e *Announcement

	e = &Announcement{
		Base:    domain.Base{},
		Title:   title,
		Content: content,
	}

	res := r.db.Create(&e)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, domain.ErrorCreateFailed
	}

	return e, nil
}

func (r *repo) Update(id string, title string, content string) (*Announcement, error) {
	if id == "" || title == "" || content == "" {
		return nil, domain.ErrorValidationNotBlank
	}

	e, err := r.Find(id)

	if err != nil {
		return nil, err
	}

	e.Title = title
	e.Content = content

	res := r.db.Save(&e)

	if res.RowsAffected == 0 {
		return e, domain.ErrorUpdateFailed
	}

	return e, nil
}

func (r *repo) Delete(id string) error {
	if id == "" {
		return domain.ErrorValidationNotBlank
	}

	res := r.db.Delete(&Announcement{}, "id = ?", id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return domain.ErrorDeleteFailed
	}

	return nil
}

func (r *repo) Paginate(page int, pageSize int, orderBy string, order string) ([]*Announcement, error) {
	var e []*Announcement

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

func (r *repo) Count() (int64, error) {
	var c int64

	var res *gorm.DB
	res = r.db.Model(&Announcement{}).Count(&c)

	if res.Error != nil {
		return 0, res.Error
	}

	return c, nil
}
