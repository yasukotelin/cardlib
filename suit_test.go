package cardlib

import (
	"testing"

	"github.com/kyokomi/emoji"
)

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

	exp := emoji.Sprint(":spades:")
	act := s.Mark()

	if exp != act {
		t.Fatalf("Faital, expected is %v actual is %v", exp, act)
	}
}
