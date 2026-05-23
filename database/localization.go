package database

import (
	"sort"

	"github.com/WindowsSov8forUs/sonolus-core-go/core"
)

type LocalizationText map[string]core.Text

func Localize(text LocalizationText, locale string, fallbackLocale string) string {
	if value, ok := text[locale]; ok {
		return string(value)
	}
	if value, ok := text[fallbackLocale]; ok {
		return string(value)
	}

	keys := make([]string, 0, len(text))
	for key := range text {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	if len(keys) == 0 {
		return ""
	}
	return string(text[keys[0]])
}
