package cardlib

import (
	"fmt"
	"strconv"
)

// Card is struct
type Card struct {
	Suit   Suit
	Number int
}

func (c *Card) String() string {
	if c.Suit == Joker {
		return fmt.Sprintf("%v%v", c.Suit.Mark(), c.Suit.String())
	}
	return fmt.Sprintf("%v%v", c.Suit.Mark(), c.GetStrNumber())
}

// GetStrNumber return card string number.
// example 13 -> "K"
func (c *Card) GetStrNumber() string {
	switch c.Number {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return strconv.Itoa(c.Number)
	}
}
