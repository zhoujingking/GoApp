package Model

type Rank int

const (
	TWO Rank = iota + 2
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

func GetRankString(rank Rank) string {
	var result string
	switch rank {
	case TWO:
		result = "2"
	case THREE:
		result = "3"
	case FOUR:
		result = "4"
	case FIVE:
		result = "5"
	case SIX:
		result = "6"
	case SEVEN:
		result = "7"
	case EIGHT:
		result = "8"
	case NINE:
		result = "9"
	case TEN:
		result = "10"
	case JACK:
		result = "J"
	case QUEEN:
		result = "Q"
	case KING:
		result = "K"
	case ACE:
		result = "A"
	}
	return result
}