package deck

import (
	"github.com/yasukotelin/cardlib/card"
	"github.com/yasukotelin/cardlib/suit"
)

// Deck is complete cards struct.
type Deck struct {
	Cards  []card.Card
	Number int
}

// NewDeck create deck.
func NewDeck() *Deck {
	var cards []card.Card
	for i := 0; i < suit.SuitNum; i++ {
		for j := 1; j <= 13; j++ {
			cards = append(cards, card.Card{
				Suit:   suit.Suit(i),
				Number: j,
			})
		}
	}
	for i := 0; i < 2; i++ {
		cards = append(cards, card.Card{
			Suit:   suit.Joker,
			Number: 0,
		})
	}
	return &Deck{
		Cards:  cards,
		Number: 54,
	}
}

func (d *Deck) RemoveJoker() {
	var rmCount int
	var cards []card.Card
	for _, c := range d.Cards {
		if c.Suit != suit.Joker {
			cards = append(cards, c)
		} else {
			rmCount++
		}
	}
	d.Cards = cards
	d.Number -= rmCount
}
