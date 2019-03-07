package deck

import (
	"math/rand"
	"time"

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
	Cards []card.Card
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
		Cards: cards,
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
		Cards: cards,
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
}

func (d *Deck) TopIndex() int {
	return 0
}

func (d *Deck) BottomIndex() int {
	return len(d.Cards) - 1
}

// Shuffle the deck. The algorythm using is Fisher–Yates shuffle
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	n := len(d.Cards)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// Draw top card from deck. Top card is 0 index of cards slice.
func (d *Deck) Draw() *card.Card {
	c := d.Cards[d.TopIndex()]
	d.Cards = d.Cards[1:len(d.Cards)]
	return &c
}
