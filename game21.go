package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cardSuit string
type cardValue 	string

const (
	spades 			cardSuit 	= 	"spade"
	hearts 			cardSuit 	= 	"heart"
	clubs 			cardSuit 	= 	"club"
	diamonds 		cardSuit 	= 	"diamond"

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

	numberOfCards	int			=	52
	numberOfSuits	int			=	4
	numberOfValues	int			=	13
)

type card struct {
	value 	cardValue
	suit 	cardSuit
	cost 	int
}

type player struct {
	cards []card
	name string
}

func (p *player) score() int {
	var score = 0
	for _, card := range p.cards {
		score += card.cost
	}
	return score
}

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

func (p *player) info() {
	fmt.Println(p.name + " cards:\n" +p.showCards())
	fmt.Println(p.name + " score:")
	fmt.Println(p.score())
}

func (p *player) takeCard(deck *[]card){
	p.cards = append(p.cards, (*deck)[len(*deck)-1])
	*deck = (*deck)[:len(*deck)-1]
}

func createDeck() []card {
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

func randomDeck(deck []card) {
	rand.Seed(time.Now().UnixNano())
	for key := range deck {
		randInt := rand.Intn(len(deck))
		deck[key], deck[randInt] = deck[randInt], deck[key]
	}
}

func computerGame(deck *[]card) player{
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

func playerGame(deck *[]card, name string) player{
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

func result(playerAI player, playerUser player, deck []card) {
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

func game(deck []card) {
	var playerAI, playerUser player
	var name string
	fmt.Println("enter your name:")
	fmt.Scanln(&name)

restart:
	for range deck {
		playerAI = computerGame(&deck)
		playerUser = playerGame(&deck, name)

		result(playerAI, playerUser, deck)

		if len(deck) == 0 {
			fmt.Println("the deck is over!")
			break
		}

		for {
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


