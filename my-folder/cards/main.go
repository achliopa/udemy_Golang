package main

<<<<<<< HEAD
import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}

=======
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
>>>>>>> 32a38f1023120a4342dfa7a63d4819aff92a3456
