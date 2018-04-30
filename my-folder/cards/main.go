package main

func main() {
	cards := newDeck()

	hand, remainignDeck := deal(cards, 5)
	hand.print()
	remainignDeck.print()
}
