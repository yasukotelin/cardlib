# cardlib

cardlib is library for playing card.

## how to use

### install

```bash
go get github.com/yasukotelin/cardlib
```

### import

```
import "github.com/yasukotelin/cardlib"
```

### simple sample code

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yasukotelin/cardlib"
)

func main() {
	// set the seed value to global random generator.
	// cardlib uses it for shuffle and cut and so on.
	rand.Seed(time.Now().UnixNano())

	// create the complete deck
	deck := cardlib.NewDeck()

	// shuffle the deck
	deck.Shuffle()

	// draw the card from Top
	card := deck.Draw()

	if card.Suit == cardlib.Joker {
		fmt.Println("I drew joker...")
	}
}
```

## example

If you would like to see this how working, there is the example main.go code in project.
You can see when Move to project root and typing `go run ./cmd/cardlib/main.go`.

```
ʕ◔ϖ◔ʔ > Make new deck!
----------------------------------------------------
♠ A ♠ 2 ♠ 3 ♠ 4 ♠ 5 ♠ 6 ♠ 7 ♠ 8 ♠ 9 ♠ 10 ♠ J ♠ Q ♠ K
♣ A ♣ 2 ♣ 3 ♣ 4 ♣ 5 ♣ 6 ♣ 7 ♣ 8 ♣ 9 ♣ 10 ♣ J ♣ Q ♣ K
♦ A ♦ 2 ♦ 3 ♦ 4 ♦ 5 ♦ 6 ♦ 7 ♦ 8 ♦ 9 ♦ 10 ♦ J ♦ Q ♦ K
♥ A ♥ 2 ♥ 3 ♥ 4 ♥ 5 ♥ 6 ♥ 7 ♥ 8 ♥ 9 ♥ 10 ♥ J ♥ Q ♥ K
🃏 Joker 🃏 Joker
----------------------------------------------------

ʕ◔ϖ◔ʔ > shuffle shuffle
----------------------------------------------------
♦ 5 ♥ K ♠ 3 ♥ 9 ♠ 5 ♠ J ♥ 8 ♥ 6 ♠ 2 ♠ 9 ♥ 3 ♦ 9 ♠ 7
♣ A ♦ J ♣ 8 ♦ K ♥ Q ♠ 8 ♦ A ♥ 10 ♣ 9 ♣ 5 ♠ Q ♠ 6 ♣ Q
♣ K ♣ J ♥ 7 ♣ 10 ♠ A ♠ 4 ♦ 7 ♥ J ♣ 4 ♥ A ♦ 3 ♣ 6 🃏 Joker
♥ 5 ♣ 3 ♦ 6 ♠ 10 ♦ Q ♣ 2 ♦ 8 ♦ 2 ♦ 4 ♠ K ♣ 7 🃏 Joker ♥ 4
♦ 10 ♥ 2
----------------------------------------------------

ʕ◔ϖ◔ʔ > Draw!!
♦ 5

ʕ◔ϖ◔ʔ > Remove joker
----------------------------------------------------
♥ K ♠ 3 ♥ 9 ♠ 5 ♠ J ♥ 8 ♥ 6 ♠ 2 ♠ 9 ♥ 3 ♦ 9 ♠ 7 ♣ A
♦ J ♣ 8 ♦ K ♥ Q ♠ 8 ♦ A ♥ 10 ♣ 9 ♣ 5 ♠ Q ♠ 6 ♣ Q ♣ K
♣ J ♥ 7 ♣ 10 ♠ A ♠ 4 ♦ 7 ♥ J ♣ 4 ♥ A ♦ 3 ♣ 6 ♥ 5 ♣ 3
♦ 6 ♠ 10 ♦ Q ♣ 2 ♦ 8 ♦ 2 ♦ 4 ♠ K ♣ 7 ♥ 4 ♦ 10 ♥ 2
----------------------------------------------------

ʕ◔ϖ◔ʔ > Draw!!
♥ K
```