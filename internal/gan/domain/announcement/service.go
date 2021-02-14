package announcement

import "github.com/v4rakh/gan/internal/gan/domain/subscription"

type Service struct {
	repo                repository
	subscriptionService *subscription.Service
}

func NewService(r repository, s *subscription.Service) *Service {
	return &Service{
		repo:                r,
		subscriptionService: s,
	}
}

func (s *Service) Get(id string) (*Announcement, error) {
	e, err := s.repo.Find(id)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (s *Service) Create(title string, content string) (Announcement, error) {
	created, err := s.repo.Create(title, content)

	if err == nil {
		s.subscriptionService.NotifySubscribers(title)
	}

	return created, err
}

func (s *Service) Update(id string, title string, content string) (*Announcement, error) {
	_, err := s.Get(id)

	if err != nil {
		return nil, err
	}

	return s.repo.Update(id, title, content)
}

func (s *Service) Delete(id string) error {
	_, err := s.Get(id)

	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *Service) Paginate(page int, pageSize int, orderBy string, order string) ([]*Announcement, error) {
	return s.repo.Paginate(page, pageSize, orderBy, order)
}

func (s *Service) Count() (int64, error) {
	return s.repo.Count()
}
