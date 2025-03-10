package main

import (
	"fmt"
	"math/rand"
)

// Initalize Constants
// Suit
const (
	Spades   = "♠"
	Hearts   = "♥"
	Diamonds = "♦"
	Clubs    = "♣"
)

// Used for Rank constants /e Iota
type Rank int

// Rank
const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Card struct
type Card struct {
	Suit string
	Rank Rank
}

// Custom Print Method (overriting)
func (c Card) String() string {
	return fmt.Sprintf("[%s - %d - %s]", c.Suit, c.Rank, c.Suit)
}

func drawCard(deck *[]Card) Card {
	card := (*deck)[0]
	*deck = (*deck)[1:]
	return card

}

func main() {

	// Assign space for the Deck as a slice
	deck := make([]Card, 0, 52)

	//  Nested loop Suits and Ranks
	//  Create an un-named card and append
	suits := []string{Spades, Hearts, Diamonds, Clubs}
	for _, suit := range suits {
		for rank := Ace; rank <= King; rank++ {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}

	// Shuffle
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	fmt.Println(deck)
	// Assign space for player hands
	playerHand := make([]Card, 0, 52)
	robotHand := make([]Card, 0, 52)

	for i := 1; i <= 14; i++ {
		switch i % 2 {
		case 0:
			playerHand = append(playerHand, drawCard(&deck))
		case 1:
			robotHand = append(robotHand, drawCard(&deck))
		}
	}

	fmt.Println(playerHand)
	fmt.Println(robotHand)

	// Draw 7 to each
}
