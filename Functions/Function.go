package Functions

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Address(s string) *string {
	caser := cases.Title(language.Und)
	upper := caser.String(s)
	return &upper
}
