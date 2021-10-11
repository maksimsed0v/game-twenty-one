package game

import (
	"fmt"

	"github.com/maksimsed0v/card/v2"
)

// player describes the player
// the struct contains the player's cards and his name
type player struct {
	name  string
	cards []card.Card
}

// score returns the sum of the player's points
func (p *player) score() (result int) {
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

// takeCard adds one card to the player from the top of the deck
func (p *player) takeCard(deck *card.Deck) {
	p.cards = append(p.cards, deck.TakeTop())
}

// info outputs information about the player's cards and score to the console
func (p *player) info() {
	fmt.Printf("%s cards:\n%s\n", p.name, p.showCards())
	fmt.Printf("%s score:\n%d\n", p.name, p.score())
}
