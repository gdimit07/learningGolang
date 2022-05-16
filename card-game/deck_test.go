package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {

		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts but got %v", d[0])
	}

	if d[len(d)-1] != "Jack of Diamonds" {
		t.Errorf("Expected Jack of Diamonds but got %v", d[len(d)-1])
	}

}

func TestNewDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
