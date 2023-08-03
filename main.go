// how do we execute this program?
// go run main.go
// the keyword 'package main' means this file will be compiled
// into an executable program (binary) called main instead of a library
package main

// short for 'format'

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}

