package main

func main() {
	cards := newDeck()

	hand, remainingcards := deal(cards, 5)
	
	hand.print()
	remainingcards.print()
}
