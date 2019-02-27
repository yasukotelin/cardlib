package main

import (
	"fmt"

	"github.com/yasukotelin/cardlib/deck"
)

func main() {
	d := deck.Make()

	for _, c := range d.Cards {
		fmt.Println(c.String(true))
	}

	fmt.Println("==========")
	fmt.Println("Shuffle")
	d.Shuffle()
	fmt.Println("==========")

	for _, c := range d.Cards {
		fmt.Println(c.String(true))
	}
}
