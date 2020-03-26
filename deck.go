package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	suits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	cards := deck{}

	for _, suit := range suits {
		for _, value := range values {
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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, d.toByteSlice(), 0666)
}
