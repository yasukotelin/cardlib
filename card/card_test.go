package card

import "testing"

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
