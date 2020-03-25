package main

import "fmt"

func main() {
	cards := []string{"Ace of Spades", "Five of Diamonds"}
	cards = append(cards, "Jack of Clubs")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
