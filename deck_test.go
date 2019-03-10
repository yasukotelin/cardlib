package cardlib

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	c := Card{
		Suit:   Spade,
		Number: 1,
	}
	if deck.Cards[0] != c {
		t.Fatalf("Faital, expected Suit is Spade, Number is 1 but actual is %v, %v", deck.Cards[0].Suit.Mark(), deck.Cards[0].Number)
	}
}

func TestNewDeckWithSuitSet(t *testing.T) {
	deck := NewDeckWithSuitSet(Heart)

	c := Card{
		Suit:   Heart,
		Number: 1,
	}
	if deck.Cards[0] != c {
		t.Fatalf("Faital, expected Suit is %v, Number is %v but actual is %v, %v", c.Suit, c.Number, deck.Cards[0].Suit, deck.Cards[0].Number)
	}
}

func TestDraw(t *testing.T) {
	deck := NewDeck()

	expNum := len(deck.Cards) - 1
	expCard := deck.Cards[deck.TopIndex()]

	c := deck.Draw()

	if len(deck.Cards) != expNum {
		t.Errorf("number of after drawing card from deck be should %d. but actualy is %d", expNum, len(deck.Cards))
	}

	if *c != expCard {
		t.Errorf("card is should be %v. but actual is %v", expCard, c)
	}
}

func TestDrawAt(t *testing.T) {
	deck := NewDeck()

	act, err := deck.DrawAt(12)
	exp := Card{
		Number: 13,
		Suit:   Spade,
	}

	if err != nil {
		t.Errorf("error should be nil. but actual is %v", err)
	}

	if *act != exp {
		t.Errorf("Draw error. expected is %v but actual %v", exp, act)
	}

	if deck.Exists(act) {
		t.Errorf("%v should be removed from the deck", *act)
	}
}

func TestExists(t *testing.T) {
	deck := NewDeck()
	deck.DrawAt(12)

	card := &Card{
		Number: 13,
		Suit:   Spade,
	}

	if deck.Exists(card) {
		t.Errorf("%v should be removed from the deck", *card)
	}

	card = &Card{
		Number: 2,
		Suit:   Heart,
	}

	if !deck.Exists(card) {
		t.Errorf("%v should be exists in the deck", *card)
	}
}

func TestRemoveJoker(t *testing.T) {
	deck := NewDeck()

	deck.RemoveJoker()

	for _, c := range deck.Cards {
		if c.Suit == Joker {
			t.Fatalf("Faital, Could not remove joker from deck.")
		}
	}
}
