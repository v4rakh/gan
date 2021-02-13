package i18n

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Service struct {
	loc *i18n.Localizer
}

func NewService() *Service {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("lang/en.json")

	return &Service{
		loc: i18n.NewLocalizer(bundle, language.English.String()),
	}
}

func (s *Service) Translate(key string, data map[string]interface{}) string {
	return s.loc.MustLocalize(&i18n.LocalizeConfig{MessageID: key, TemplateData: data})
}
