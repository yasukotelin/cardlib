package suit

import "testing"

func TestSuitString(t *testing.T) {
	s := Spade

	exp := "Spade"
	act := s.String()

	if exp != act {
		t.Fatalf("Faital, expected is %v actual is %v", exp, act)
	}
}

func TestSuitMark(t *testing.T) {
	s := Spade

	exp := "♠"
	act := s.Mark()

	if exp != act {
		t.Fatalf("Faital, expected is %v actual is %v", exp, act)
	}
}
