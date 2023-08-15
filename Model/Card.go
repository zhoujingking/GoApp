package Model

import (
	"strings"
)

type Card struct {
	Rank Rank
	Suit Suit
}

func (card Card) Display() string {
	return strings.Join([]string{GetRankString(card.Rank), "of", GetSuitString(card.Suit)}, " ")
}