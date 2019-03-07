package cardlib

import (
	"testing"
)

func TestMake(t *testing.T) {
	deck := Make()

	c := Card{
		Suit:   Spade,
		Number: 1,
	}
	if deck.Cards[0] != c {
		t.Fatalf("Faital, expected Suit is Spade, Number is 1 but actual is %v, %v", deck.Cards[0].Suit.Mark(), deck.Cards[0].Number)
	}
}

func TestMakeSuitSet(t *testing.T) {
	deck := MakeSuitSet(Heart)

	c := Card{
		Suit:   Heart,
		Number: 1,
	}
	if deck.Cards[0] != c {
		t.Fatalf("Faital, expected Suit is %v, Number is %v but actual is %v, %v", c.Suit, c.Number, deck.Cards[0].Suit, deck.Cards[0].Number)
	}
}

func TestRemoveJoker1(t *testing.T) {
	deck := Make()

	deck.RemoveJoker()

	for _, c := range deck.Cards {
		if c.Suit == Joker {
			t.Fatalf("Faital, Could not remove joker from deck.")
		}
	}
}

func TestDraw(t *testing.T) {
	deck := Make()

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
