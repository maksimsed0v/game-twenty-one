package card

// Card describes the card.
// the struct contains the values and suits of the card
type Card struct {
	Value 	Value
	Suit 	Suit
}

// Suit - constants for card suits
type Suit string
const (
	Spades   Suit = 	"spade"
	Hearts   Suit = 	"heart"
	Clubs    Suit = 	"club"
	Diamonds Suit = 	"diamond"
)

// Value - constants for card values
type Value string
const (
	Two   Value =	"2"
	Three Value = 	"3"
	Four  Value = 	"4"
	Five  Value = 	"5"
	Six   Value =	"6"
	Seven Value =	"7"
	Eight Value =	"8"
	Nine  Value =	"9"
	Ten   Value =	"10"
	Jack  Value =	"J"
	Queen Value =	"Q"
	King  Value =	"K"
	Ace   Value =	"A"
)