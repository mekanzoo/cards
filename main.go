package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())
	deck := newDeck()
	deck.shuffle()
	deck.print()
	// cards.saveToFile("halo")
}

func newCard() string {
	return "Five of Diamonds"
}
