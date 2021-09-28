package card

import (
	"math/rand"
	"time"
)

// Card describes the card.
// the struct contains the values and suits of the card
type Card struct {
	Value Value
	Suit  Suit
}

// Suit - constants for card suits
type Suit string

const (
	Spades   Suit = "spade"
	Hearts   Suit = "heart"
	Clubs    Suit = "club"
	Diamonds Suit = "diamond"
)

// Value - constants for card values
type Value string

const (
	Two   Value = "2"
	Three Value = "3"
	Four  Value = "4"
	Five  Value = "5"
	Six   Value = "6"
	Seven Value = "7"
	Eight Value = "8"
	Nine  Value = "9"
	Ten   Value = "10"
	Jack  Value = "J"
	Queen Value = "Q"
	King  Value = "K"
	Ace   Value = "A"
)

const (
	// NumberOfSuits - length of the suits array
	NumberOfSuits int = 4

	// NumberOfValues - length of the values array
	NumberOfValues int = 13
)

// CreateDeck generates a deck of cards and returns it as a slice
func CreateDeck() []Card {
	// suits - an array of all suits of cards
	suits := [NumberOfSuits]Suit{Spades, Hearts, Clubs, Diamonds}

	// values - an array of all values of cards
	values := [NumberOfValues]Value{Two, Three, Four, Five, Six,
		Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	// deck - a slice of a deck of cards
	deck := make([]Card, NumberOfValues*NumberOfSuits)

	// current - a variable for the current card in the deck
	current := 0

	for v := 0; v < NumberOfValues; v++ {
		for s := 0; s < NumberOfSuits; s++ {
			deck[current].Value = values[v]
			deck[current].Suit = suits[s]
			current += 1
		}
	}
	return deck
}

// RandomDeck shuffles the deck of cards
func RandomDeck(deck []Card) {
	// function for randomness
	rand.Seed(time.Now().UnixNano())

	for key := range deck {
		randInt := rand.Intn(len(deck))
		deck[key], deck[randInt] = deck[randInt], deck[key]
	}
}
