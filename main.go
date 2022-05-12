package main

import "fmt"

// func newCard() string {
// 	return "Five of Diamonds"
// }

func main() {

	mysliuce := []string{"hello"}

	fmt.Println(mysliuce)

	//cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	//cards := deck{newCard(), "Ace of Diamonds"}
	//cards = append(cards, "Six of hearts")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards := newDeck()
	cards.print()

}
