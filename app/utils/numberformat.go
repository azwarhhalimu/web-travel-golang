package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func NumberFormat(nilai int) string {
	p := message.NewPrinter(language.Indonesian)
	currency := p.Sprintf("Rp. %d", nilai)
	return currency
}
