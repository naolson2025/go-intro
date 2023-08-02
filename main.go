// how do we execute this program?
// go run main.go
// the keyword 'package main' means this file will be compiled
// into an executable program (binary) called main instead of a library
package main

// short for 'format'
import "fmt"

func main() {
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")
	
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
