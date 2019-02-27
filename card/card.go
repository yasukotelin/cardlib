package card

import (
	"fmt"
	"strconv"

	"github.com/yasukotelin/cardlib/suit"
)

// Card is struct
type Card struct {
	Suit   suit.Suit
	Number int
}

func (c *Card) String(mark bool) string {
	if c.Suit == suit.Joker {
		return c.Suit.Mark()
	}
	if mark {
		return fmt.Sprintf("%v%v", c.Suit.Mark(), c.GetStrNumber())
	}
	return fmt.Sprintf("(%v)%v", c.Suit.String(), c.GetStrNumber())
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
