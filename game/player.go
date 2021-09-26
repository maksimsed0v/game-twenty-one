package game

import (
	"fmt"

	"../card"
)

// player describes the player
// the struct contains the player's cards and his name
type player struct {
	cards []card.Card
	name  string
}

// score returns the sum of the player's points
func (p *player) score() (result int) {
	costs := map[card.Value]int{
		card.Two:   2,
		card.Three: 3,
		card.Four:  4,
		card.Five:  5,
		card.Six:   6,
		card.Seven: 7,
		card.Eight: 8,
		card.Nine:  9,
		card.Ten:   10,
		card.Jack:  2,
		card.Queen: 3,
		card.King:  4,
		card.Ace:   11,
	}
	for _, c := range p.cards {
		result += costs[c.Value]
	}
	return
}

// showCards returns the values and suits of all the player's cards as a string
func (p *player) showCards() (allCards string) {
	for _, c := range p.cards {
		if c.Value != "" {
			allCards += fmt.Sprintf("%s %s, ", c.Value, c.Suit)
		}
	}
	if allCards == "" {
		return "no cards"
	}
	return allCards[:len(allCards)-2] + "."
}

// info outputs information about the player's cards and score to the console
func (p *player) info() {
	fmt.Println(p.name + " cards:\n" + p.showCards())
	fmt.Println(p.name + " score:")
	fmt.Println(p.score())
}

// takeCard adds one card to the player from the top of the deck
func (p *player) takeCard(deck *[]card.Card) {
	p.cards = append(p.cards, (*deck)[len(*deck)-1])
	*deck = (*deck)[:len(*deck)-1]
}
