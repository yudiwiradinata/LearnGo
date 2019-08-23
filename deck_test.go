package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Excpect 16, but %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Excpect Ace of Spades, but %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Excpect Four of Clubs, but %v", d[len(d)-1])
	}
}
