package main

import "fmt"

// How to use functions as classes in Go
// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// by declare a fn that takes a 'deck' all decks can now access that fn
// example: deck.print()
// Loop through our deck of cards and print each one out
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}