package deck

import (
	"testing"

	"github.com/yasukotelin/cardlib/card"
	"github.com/yasukotelin/cardlib/suit"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	if deck.Number != len(deck.Cards) {
		t.Fatalf("Faital, deck.Number is %v but len(deck.Cards) is %v", deck.Number, len(deck.Cards))
	}

	c := card.Card{
		Suit:   suit.Spade,
		Number: 1,
	}
	if deck.Cards[0] != c {
		t.Fatalf("Faital, expected Suit is Spade, Number is 1 but actual is %v, %v", deck.Cards[0].Suit, deck.Cards[0].Number)
	}
}

func TestRemoveJoker1(t *testing.T) {
	deck := NewDeck()

	deck.RemoveJoker()

	for _, c := range deck.Cards {
		if c.Suit == suit.Joker {
			t.Fatalf("Faital, Could not remove joker from deck.")
		}
	}
}

func TestRemoveJoker2(t *testing.T) {
	deck := NewDeck()

	exp := deck.Number - 2

	deck.RemoveJoker()
	act := deck.Number

	if exp != act {
		t.Fatalf("Faital, expected is %v actual is %v", exp, act)
	}
}
