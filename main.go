package main

func main() {
	cards := newDeck()
	cards = append(cards, "Joker")
	hand, remainingDeck := deal(cards, 5)
	hand.saveToFile("my_hand")
	remainingDeck.saveToFile("my_remainingdeck")
}
