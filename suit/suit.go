package suit

type Suit int

const (
	Spade Suit = iota
	Club
	Diamond
	Heart
	Joker

	// JokerNum is defined number for joker
	JokerNum = 14
)

func (s Suit) String() string {
	switch s {
	case Spade:
		return "Spade"
	case Club:
		return "Club"
	case Diamond:
		return "Diamond"
	case Heart:
		return "Heart"
	case Joker:
		return "Joker"
	default:
		return ""
	}
}

// Mark return Suit mark by Unicode Emoji
func (s Suit) Mark() string {
	switch s {
	case Spade:
		return "♠"
	case Club:
		return "♣"
	case Diamond:
		return "♦"
	case Heart:
		return "♥"
	case Joker:
		return "Joker"
	default:
		return ""
	}
}
