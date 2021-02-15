package i18n

import (
	"bytes"
	"encoding/json"
	"github.com/markbates/pkger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

type Service struct {
	loc *i18n.Localizer
}

func NewService() *Service {
	const fileNameEnglish = "/lang/en.json"
	langFile, err := pkger.Open(fileNameEnglish)

	if err != nil {
		log.Fatalf("Could not open file '%s'. Reason: %s", fileNameEnglish, err)
	}
	defer langFile.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(langFile)

	if err != nil {
		log.Fatalf("Could not read file contents of '%s'. Reason: %s", fileNameEnglish, err)
	}

	log.Printf("Loaded file contents of '%s'", fileNameEnglish)

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustParseMessageFileBytes(buf.Bytes(), "en.json")

	return &Service{
		loc: i18n.NewLocalizer(bundle, language.English.String()),
	}
}

func (s *Service) Translate(key string, data map[string]interface{}) string {
	return s.loc.MustLocalize(&i18n.LocalizeConfig{MessageID: key, TemplateData: data})
}
