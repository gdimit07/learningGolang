package main

import "fmt"

// func newCard() string {
// 	return "Five of Diamonds"
// }

func main() {

	//cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	//cards := deck{newCard(), "Ace of Diamonds"}
	//cards = append(cards, "Six of hearts")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()

	remainingCards.print()

	fmt.Println(cards.toString())

}
