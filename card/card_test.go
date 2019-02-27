package card

import (
	"testing"

	"github.com/yasukotelin/cardlib/suit"
)

func TestString1(t *testing.T) {
	c := Card{
		Number: 1,
		Suit:   suit.Diamond,
	}

	exp := "♦A"
	act := c.String()

	if exp != act {
		t.Fatalf("Fatail, expected is %s actual is %s", exp, act)
	}
}

func TestString2(t *testing.T) {
	c := Card{
		Number: 10,
		Suit:   suit.Heart,
	}

	exp := "♥10"
	act := c.String()

	if exp != act {
		t.Fatalf("Fatail, expected is %s actual is %s", exp, act)
	}
}

func TestString3(t *testing.T) {
	c := Card{
		Number: suit.JokerNum,
		Suit:   suit.Joker,
	}

	exp := "Joker"
	act := c.String()

	if exp != act {
		t.Fatalf("Fatail, expected is %s actual is %s", exp, act)
	}
}

func TestGetStrNumber1(t *testing.T) {
	c := Card{
		Number: 10,
	}

	exp := "10"
	act := c.GetStrNumber()

	if exp != act {
		t.Fatalf("Fatail, expected is %s actual is %s", exp, act)
	}
}

func TestGetStrNumber2(t *testing.T) {
	c := Card{
		Number: 11,
	}

	exp := "J"
	act := c.GetStrNumber()

	if exp != act {
		t.Fatalf("Fatail, expected is %s actual is %s", exp, act)
	}
}

func TestGetStrNumber3(t *testing.T) {
	c := Card{
		Number: 1,
	}

	exp := "A"
	act := c.GetStrNumber()

	if exp != act {
		t.Fatalf("Fatail, expected is %s actual is %s", exp, act)
	}
}
