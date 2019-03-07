package cardlib

import (
	"math/rand"
	"time"
)

const (
	suitSetNum  = 13
	jokerNum    = 2
	suitKindNum = 4
)

// Deck is card set struct.
type Deck struct {
	Cards []Card
}

// Make returns complete new deck
func Make() *Deck {
	var cards []Card
	for i := 0; i < suitKindNum; i++ {
		for j := 1; j <= suitSetNum; j++ {
			cards = append(cards, Card{
				Suit:   Suit(i),
				Number: j,
			})
		}
	}
	for i := 0; i < jokerNum; i++ {
		cards = append(cards, Card{
			Suit:   Joker,
			Number: 0,
		})
	}
	return &Deck{
		Cards: cards,
	}
}

// MakeSuitSet returns new deck maked only suit of arg
func MakeSuitSet(suit Suit) *Deck {
	var cards []Card
	for i := 1; i <= suitSetNum; i++ {
		cards = append(cards, Card{
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
	var cards []Card
	for _, c := range d.Cards {
		if c.Suit != Joker {
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
func (d *Deck) Draw() *Card {
	c := d.Cards[d.TopIndex()]
	d.Cards = d.Cards[1:len(d.Cards)]
	return &c
}
