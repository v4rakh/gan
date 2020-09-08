package announcement

import (
	"github.com/v4rakh/gan/internal/gan/domain"
	"gorm.io/gorm"
)

type repository interface {
	List(orderBy string, order string) ([]*Announcement, error)
	Find(id string) (*Announcement, error)
	Create(title string, content string) (Announcement, error)
	Update(id string, title string, content string) (*Announcement, error)
	Delete(id string) error
}

type sqliteRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *sqliteRepo {
	return &sqliteRepo{
		db: db,
	}
}

func (r *sqliteRepo) List(orderBy string, order string) ([]*Announcement, error) {
	var e []*Announcement

	var res *gorm.DB
	if orderBy != "" && order != "" {
		res = r.db.Order(orderBy + " " + order).Find(&e)
	} else {
		res = r.db.Find(&e)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return e, nil
}

func (r *sqliteRepo) Find(id string) (*Announcement, error) {
	var announcement Announcement

	res := r.db.Find(&announcement, "id = ?", id)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, domain.ErrorAnnouncementNotFound
	}

	return &announcement, nil
}

func (r *sqliteRepo) Create(title string, content string) (Announcement, error) {
	var e Announcement
	e = Announcement{
		Base:    domain.Base{},
		Title:   title,
		Content: content,
	}

	res := r.db.Create(&e)

	if res.Error != nil {
		return e, res.Error
	}

	if res.RowsAffected == 0 {
		return e, domain.ErrorAnnouncementCreateFailed
	}

	return e, nil
}

func (r *sqliteRepo) Update(id string, title string, content string) (*Announcement, error) {
	e, err := r.Find(id)

	if err != nil {
		return nil, err
	}

	e.Title = title
	e.Content = content

	res := r.db.Save(&e)

	if res.RowsAffected == 0 {
		return e, domain.ErrorAnnouncementUpdateFailed
	}

	return e, nil
}

func (r *sqliteRepo) Delete(id string) error {
	res := r.db.Delete(&Announcement{}, "id = ?", id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return domain.ErrorAnnouncementDeleteFailed
	}

	return nil
}
