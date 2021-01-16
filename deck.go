package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cardType := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, a := range cardType {
		for _, b := range cardNumber {
			cards = append(cards, a+" of "+fmt.Sprintf("%d", b))
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, "|", card)
	}
}
