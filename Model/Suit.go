package Model

type Suit int

const (
	SPADE Suit = iota + 1
	HEART
	DIAMOND
	CLUB
)

func GetSuitString(suit Suit) string {
	var result string
	switch suit {
	case SPADE:
		result = "Spade"
	case HEART:
		result = "Heart"
	case DIAMOND:
		result = "Diamond"
	case CLUB:
		result = "Club"
	
	}
	return result
}