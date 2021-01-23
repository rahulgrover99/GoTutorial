package main

func main() {
	// cards := newDeck()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("mycards")
	cards := newDeckFromFile("mycards")
	cards.shuffle()
	cards.print()
}
