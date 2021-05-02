package subscription

import (
	"github.com/v4rakh/gan/internal/errors"
	"github.com/v4rakh/gan/internal/gan/constant"
	"github.com/v4rakh/gan/internal/gan/service/i18n"
	"github.com/v4rakh/gan/internal/gan/service/mail"
	"github.com/v4rakh/gan/internal/util"
	"log"
	"net/url"
	"os"
)

const randomTokenLength = 20

type Service struct {
	repo        repository
	mailService *mail.Service
	i18nService *i18n.Service
}

func NewService(r repository, m *mail.Service, i *i18n.Service) *Service {
	return &Service{
		repo:        r,
		mailService: m,
		i18nService: i,
	}
}

func (s *Service) Get(address string) (*Subscription, error) {
	if address == "" {
		return nil, errors.ErrorValidationNotBlank
	}

	e, err := s.repo.Find(address)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (s *Service) Rescue(address string) error {
	if address == "" {
		return errors.ErrorValidationNotBlank
	}

	e, err := s.Get(address)
	if err != nil {
		return err
	}

	switch State(e.State) {
	case Pending:
		s.mailService.Send(address, s.i18nService.Translate("mail_subscription_created_subject", nil), s.i18nService.Translate("mail_subscription_created_body", map[string]interface{}{
			"Domain":  os.Getenv(constant.EnvDomain),
			"Address": url.QueryEscape(address),
			"Token":   url.QueryEscape(e.Token),
		}))
	case Active:
		s.mailService.Send(address, s.i18nService.Translate("mail_subscription_verified_subject", nil), s.i18nService.Translate("mail_subscription_verified_body", map[string]interface{}{
			"Domain":  os.Getenv(constant.EnvDomain),
			"Address": url.QueryEscape(address),
			"Token":   url.QueryEscape(e.Token),
		}))
	}

	return nil
}

func (s *Service) Create(address string) error {
	if address == "" {
		return errors.ErrorValidationNotBlank
	}

	e, err := s.Get(address)
	if e != nil {
		return errors.ErrorSubscriptionAlreadyActive
	}

	token := util.RandomString(randomTokenLength)
	err = s.repo.Create(address, Pending, token)
	s.mailService.Send(address, s.i18nService.Translate("mail_subscription_created_subject", nil), s.i18nService.Translate("mail_subscription_created_body", map[string]interface{}{
		"Domain":  os.Getenv(constant.EnvDomain),
		"Address": url.QueryEscape(address),
		"Token":   url.QueryEscape(token),
	}))
	return err
}

func (s *Service) Verify(address string, token string) error {
	if address == "" || token == "" {
		return errors.ErrorValidationNotBlank
	}

	e, err := s.Get(address)
	if err != nil {
		return err
	}

	if e.Token != token {
		return errors.ErrorSubscriptionForbiddenTokenMatch
	}

	if State(e.State) == Active {
		return errors.ErrorSubscriptionAlreadyActive
	}

	newToken := util.RandomString(randomTokenLength)
	updated, err := s.repo.Update(address, Active, newToken)

	if updated != nil {
		s.mailService.Send(address, s.i18nService.Translate("mail_subscription_verified_subject", nil), s.i18nService.Translate("mail_subscription_verified_body", map[string]interface{}{
			"Domain":  os.Getenv(constant.EnvDomain),
			"Address": url.QueryEscape(address),
			"Token":   url.QueryEscape(newToken),
		}))
	}

	return err
}

func (s *Service) Delete(address string, token string) error {
	if address == "" || token == "" {
		return errors.ErrorValidationNotBlank
	}

	e, err := s.Get(address)
	if err != nil {
		return err
	}

	if e.Token != token {
		return errors.ErrorSubscriptionForbiddenTokenMatch
	}

	err = s.repo.Delete(e.Address)
	if err != nil {
		return err
	}

	s.mailService.Send(address, s.i18nService.Translate("mail_subscription_deleted_subject", nil), s.i18nService.Translate("mail_subscription_deleted_body", map[string]interface{}{
		"Domain": os.Getenv(constant.EnvDomain),
	}))

	return nil
}

func (s *Service) NotifySubscribers(title string) {
	if title == "" {
		return
	}

	e, err := s.repo.ListWhereState(Active)
	if err != nil {
		log.Printf("Could not retrieve subscriptions. Reason: %s\n", err.Error())
		return
	}

	if len(e) == 0 {
		log.Print("No active subscriptions found, skipping notifications")
		return
	}

	for _, sub := range e {
		s.mailService.Send(sub.Address, s.i18nService.Translate("mail_new_announcement_subject", nil), s.i18nService.Translate("mail_new_announcement_body", map[string]interface{}{
			"Domain": os.Getenv(constant.EnvDomain),
			"Title":  title,
		}))
	}
}

func (s *Service) DeleteByAddress(address string) error {
	if address == "" {
		return errors.ErrorValidationNotBlank
	}

	_, err := s.Get(address)
	if err != nil {
		return err
	}

	return s.repo.Delete(address)
}

func (s *Service) Paginate(page int, pageSize int, orderBy string, order string) ([]*Subscription, error) {
	return s.repo.Paginate(page, pageSize, orderBy, order)
}

func (s *Service) Count() (int64, error) {
	return s.repo.Count()
}
