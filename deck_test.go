package main

import "testing"

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()
	if len(testDeck) != 52 {
		t.Errorf("Expected deck to have a size of 52, but instead found %v.", len(testDeck))
	}

	if testDeck[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be Ace of Hearts, but instead found, %v", testDeck[0])
	}

	if testDeck[len(testDeck)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but instead found, %v", testDeck[len(testDeck)-1])
	}
}
