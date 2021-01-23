package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(cards))
	}
	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %s", cards[0])
	}
	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %s", cards[len(cards)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
