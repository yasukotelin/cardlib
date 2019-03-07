package cardlib

import (
	"github.com/kyokomi/emoji"
)

type Suit int

const (
	Spade Suit = iota
	Club
	Diamond
	Heart
	Joker
)

func (s *Suit) String() string {
	switch *s {
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
func (s *Suit) Mark() string {
	switch *s {
	case Spade:
		return emoji.Sprint(":spades:")
	case Club:
		return emoji.Sprint(":clubs:")
	case Diamond:
		return emoji.Sprint(":diamonds:")
	case Heart:
		return emoji.Sprint(":hearts:")
	case Joker:
		return emoji.Sprint(":black_joker:")
	default:
		return ""
	}
}
