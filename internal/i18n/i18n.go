package i18n

import (
	"embed"
	"fmt"

	"github.com/retalkgo/retalk/internal/config"
	"gopkg.in/yaml.v3"
)

//go:embed translations/*
var translationFiles embed.FS

type I18nSchema struct {
	LanguageCode string
	Translations map[string]string
}

var i18nInterface *I18nSchema

func initI18n(languageCode string) {
	translations := make(map[string]string)
	filePath := fmt.Sprintf("translations/%s.yml", languageCode)

	content, err := translationFiles.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &translations)
	if err != nil {
		panic(err)
	}

	i18nInterface = &I18nSchema{
		LanguageCode: languageCode,
		Translations: translations,
	}
}

func I18n(key string) string {
	if i18nInterface == nil {
		lang := config.Config().Lang
		initI18n(lang)
	}
	return i18nInterface.Translations[key]
}
