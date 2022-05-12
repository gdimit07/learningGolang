package main

import "fmt"

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Hearts", "Clubs", "Spades", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
