package deck

import (
	"github.com/yasukotelin/cardlib/card"
	"github.com/yasukotelin/cardlib/suit"
)

const (
	// SuitSetNum is max number of suit set
	SuitSetNum = 13
	// SuitKindNum is number of kind of suit
	SuitKindNum = 4
	// JokerNum is max number of joker in deck
	JokerNum = 2
)

// Deck is card set struct.
type Deck struct {
	Cards  []card.Card
	Number int
}

// Make returns complete new deck
func Make() *Deck {
	var cards []card.Card
	for i := 0; i < SuitKindNum; i++ {
		for j := 1; j <= SuitSetNum; j++ {
			cards = append(cards, card.Card{
				Suit:   suit.Suit(i),
				Number: j,
			})
		}
	}
	for i := 0; i < JokerNum; i++ {
		cards = append(cards, card.Card{
			Suit:   suit.Joker,
			Number: suit.JokerNum,
		})
	}
	return &Deck{
		Cards:  cards,
		Number: 54,
	}
}

// MakeSuitSet returns new deck maked only suit of arg
func MakeSuitSet(suit suit.Suit) *Deck {
	var cards []card.Card
	for i := 1; i <= SuitSetNum; i++ {
		cards = append(cards, card.Card{
			Suit:   suit,
			Number: i,
		})
	}
	return &Deck{
		Cards:  cards,
		Number: SuitSetNum,
	}
}

// RemoveJoker is remove all joker from deck
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
