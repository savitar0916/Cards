package main

import "fmt"

func main() {
	// card := "Ace of Spades"
	// fmt.Println(card)
	// card = newCard()
	// fmt.Println(card)

	// cards := deck{newCard(), newCard()}
	// cards = append(cards, "Six of Spades")

	// cards := newDeck()
	// cards.print()

	// hand, remainingCards := deal(cards, 7)

	// fmt.Println("First")
	// hand.print()
	// fmt.Println("Second")
	// remainingCards.print()

	// greeting := "Hello World!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_card")

	cards := newDeckFromFile("my_card")
	cards.print()
	fmt.Println("AFTER")
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
