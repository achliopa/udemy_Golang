package main

// import (
// 	"fmt"
// )

func main() {
	// cards := newDeck()

	// hand, remainignDeck := deal(cards, 5)
	// hand.print()
	// remainignDeck.print()

	// test type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// test string join
	// cards := newDeck()
	// fmt.Println(cards.toString())

	//test saveToFile()
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// test newDeckFromFile
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// test shuffle
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
