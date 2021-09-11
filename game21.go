package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cardSuit string
type cardValue 	string

const (
	// cardSuit - constants for card suits
	spades 			cardSuit 	= 	"spade"
	hearts 			cardSuit 	= 	"heart"
	clubs 			cardSuit 	= 	"club"
	diamonds 		cardSuit 	= 	"diamond"

	// cardValue - constants for card values
	two 			cardValue 	=	"2"
	three			cardValue 	= 	"3"
	four 			cardValue 	= 	"4"
	five 			cardValue 	= 	"5"
	six				cardValue	=	"6"
	seven			cardValue	=	"7"
	eight			cardValue	=	"8"
	nine			cardValue	=	"9"
	ten				cardValue	=	"10"
	jack			cardValue	=	"J"
	queen			cardValue	=	"Q"
	king			cardValue	=	"K"
	ace				cardValue	=	"A"

	// constants for the length of arrays
	numberOfCards	int			=	52
	numberOfSuits	int			=	4
	numberOfValues	int			=	13
)

// the card struct describes the card.
// the struct contains the values, suits and costs of the card
type card struct {
	value 	cardValue
	suit 	cardSuit
	cost 	int
}

// the player struct describes the player
// the struct contains the player's cards and his name
type player struct {
	cards []card
	name string
}

// the score method returns the sum of the player's points
func (p *player) score() int {
	var score = 0
	for _, card := range p.cards {
		score += card.cost
	}
	return score
}

// the showCards method returns the values and suits of all the player's cards as a string
func (p *player) showCards() string {
	var allCards string
	for _, card := range p.cards {
		if card.value != ""{
			allCards += string(card.value) + " " + string(card.suit) + ", "
		}
	}
	if allCards == "" {
		return "no cards"
	}
	return allCards[:len(allCards) - 2] + "."
}

// the info method outputs information about the player's cards and score to the console
func (p *player) info() {
	fmt.Println(p.name + " cards:\n" +p.showCards())
	fmt.Println(p.name + " score:")
	fmt.Println(p.score())
}

// the takeCard method adds one card to the player from the top of the deck
// input parameter: deck of cards
func (p *player) takeCard(deck *[]card) {
	p.cards = append(p.cards, (*deck)[len(*deck)-1])
	*deck = (*deck)[:len(*deck)-1]
}

// func createDeck generates a deck of cards and returns it as a slice
func createDeck() []card {
	// suits - an array of all suits of cards
	// values - an array of all values of cards
	// costs - an array of all costs of cards
	// deck - a slice of a deck of cards
	// current - a variable for the current card in the deck
	suits := [numberOfSuits]cardSuit{spades, hearts, clubs, diamonds}
	values := [numberOfValues]cardValue{two, three, four, five, six, seven, eight, nine, ten, jack, queen, king, ace}
	costs := [numberOfValues]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 2, 3, 4, 11}
	deck := make([]card, numberOfCards)
	var current = 0
	for v := 0; v < numberOfValues ; v++ {
		for s := 0; s < numberOfSuits ; s++ {
			deck[current].value = values[v]
			deck[current].suit = suits[s]
			deck[current].cost = costs[v]
			current += 1
		}
	}
	return deck
}

// func randomDeck shuffles the deck of cards
// input parameter: deck of cards
func randomDeck(deck []card) {
	rand.Seed(time.Now().UnixNano())
	for key := range deck {
		randInt := rand.Intn(len(deck))
		deck[key], deck[randInt] = deck[randInt], deck[key]
	}
}

// func computerGame creates a computer player and gives him cards
// input parameter: deck of cards
// output parameter: computer player
func computerGame(deck *[]card) player{
	// playerAI - the object of the player structure
	var playerAI player
	playerAI.name = "computer"
	for range *deck {
		if playerAI.score() < 17 {
			playerAI.takeCard(deck)
		} else {
			break
		}
	}
	return playerAI
}

// func playerGame creates a user player and gives him cards
// input parameters: deck of cards, name of user
// output parameter: user player
func playerGame(deck *[]card, name string) player{
	// playerUser - the object of the player structure
	var playerUser player
	playerUser.name = name

	if len(*deck) != 0 {
		playerUser.takeCard(deck)
		playerUser.info()
	} else {
		return playerUser
	}

stop:
	for range *deck {
		// answer - a variable for the player's response
		var answer string
		fmt.Println("more? (Y/N)")
		fmt.Scanln(&answer)
		switch answer{
		case "Y", "y":
			playerUser.takeCard(deck)
			playerUser.info()
		case "N", "n":
			break stop
		default:
			fmt.Println("invalid response, try again")
			continue stop
		}
	}
	return playerUser
}

// func result outputs the result of the round to the console
// input parameters: computer player, user player
func result(playerAI player, playerUser player) {
	fmt.Println("------------------------------------")
	fmt.Println("---------------RESULT---------------")
	fmt.Println("------------------------------------")

	playerAI.info()
	playerUser.info()

	if playerAI.score() > 21 && playerUser.score() > 21 {
		fmt.Println("everyone lost!")
	} else if (playerAI.score() > playerUser.score() && playerAI.score() <= 21) || (playerUser.score() > 21) {
		fmt.Println(playerAI.name + " won!")
	} else if (playerAI.score() < playerUser.score() && playerUser.score() <= 21) || (playerAI.score() > 21) {
		fmt.Println(playerUser.name + " won!")
	} else if playerAI.score() == playerUser.score() {
		fmt.Println("draw!")
	}
}

// func game performs the main function of the game
// input parameter: deck of cards
func game(deck []card) {
	// playerAI, playerUser - the object of the player structure
	// name - name of user
	var playerAI, playerUser player
	var name string
	fmt.Println("enter your name:")
	fmt.Scanln(&name)

restart:
	for range deck {
		playerAI = computerGame(&deck)
		playerUser = playerGame(&deck, name)

		result(playerAI, playerUser)

		if len(deck) == 0 {
			fmt.Println("the deck is over!")
			break
		}

		for {
			// answer - a variable for the player's response
			var answer string
			fmt.Println("start the game again? (Y/N)")
			fmt.Scanln(&answer)
			switch answer {
			case "Y", "y":
				continue restart
			case "N", "n":
				break restart
			default:
				fmt.Println("invalid response, try again")
				continue
			}
		}
	}
}

func main() {

	deck := createDeck()
	randomDeck(deck)

	game(deck)
}


