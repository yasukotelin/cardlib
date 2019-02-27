package main

import (
	"fmt"

	"github.com/yasukotelin/cardlib/card"
	"github.com/yasukotelin/cardlib/deck"
)

func main() {
	fmt.Println("\nʕ◔ϖ◔ʔ > Make new deck!")
	d := deck.Make()
	printlnAllCards(&d.Cards)

	fmt.Println("\nʕ◔ϖ◔ʔ > shuffle shuffle")
	d.Shuffle()
	printlnAllCards(&d.Cards)
}

func printlnAllCards(cards *[]card.Card) {
	fmt.Println("----------------------------------------------------")
	for i, c := range *cards {
		fmt.Print(c.String())
		if (i+1)%deck.SuitSetNum == 0 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	fmt.Println("----------------------------------------------------")
}
