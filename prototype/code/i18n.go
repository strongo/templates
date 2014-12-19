package code

import (
	"github.com/strongo/templates"
)

type I18nInCode struct {
	messages map[string]string
}

func CreateMessages() map[string]I18nInCode {
	return map[string]I18nInCode{
		"ru_RU": I18nInCode{
			messages: map[string]string{
				"Welcome to page": "Добро пожаловать на страницу",
				"Put your content here.": "Разместите ваш контент здесь",
			},
		},
	}
}

var I18Storage map[string]I18nInCode = CreateMessages()

func GetI18N(locale string) templates.I18n {
	return templates.I18n(I18Storage[locale])
}

func (i18n I18nInCode) GetText(key string) string {
	return i18n.messages[key]
}
