package main

import (
	"fmt"

	"github.com/yasukotelin/cardlib"
)

const (
	gopher = "ʕ◔ϖ◔ʔ"
)

func main() {
	fmt.Printf("\n%v > Make new deck!\n", gopher)
	d := cardlib.NewDeck()
	printlnAllCards(&d.Cards)

	fmt.Printf("\n%v > shuffle shuffle\n", gopher)
	d.Shuffle()
	printlnAllCards(&d.Cards)

	fmt.Printf("\n%v > Draw!!\n", gopher)
	c := d.Draw()
	fmt.Println(c.String())

	fmt.Printf("\n%v > Remove joker\n", gopher)
	d.RemoveJoker()
	printlnAllCards(&d.Cards)

	fmt.Printf("\n%v > Draw!!\n", gopher)
	c = d.Draw()
	fmt.Println(c.String())
}

func printlnAllCards(cards *[]cardlib.Card) {
	fmt.Println("----------------------------------------------------")
	for i, c := range *cards {
		fmt.Print(c.String())
		if (i+1)%13 == 0 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	fmt.Println("----------------------------------------------------")
}
