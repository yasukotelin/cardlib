package cardlib

import (
	"testing"

	"github.com/kyokomi/emoji"
)

func TestString1(t *testing.T) {
	c := Card{
		Number: 1,
		Suit:   Diamond,
	}

	exp := emoji.Sprint(":diamonds:A")
	act := c.String()

	if exp != act {
		t.Fatalf("Fatail, expected is %s actual is %s", exp, act)
	}
}

func TestString2(t *testing.T) {
	c := Card{
		Number: 10,
		Suit:   Heart,
	}

	exp := emoji.Sprint(":hearts:10")
	act := c.String()

	if exp != act {
		t.Fatalf("Fatail, expected is %s actual is %s", exp, act)
	}
}

func TestString3(t *testing.T) {
	c := Card{
		Number: 0,
		Suit:   Joker,
	}

	exp := emoji.Sprint(":black_joker:Joker")
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
