package main

import (
	"fmt"
	"math/rand"
	"time"
)

// cardSuit - constants for card suits
type cardSuit string
const (
	spades 			cardSuit 	= 	"spade"
	hearts 			cardSuit 	= 	"heart"
	clubs 			cardSuit 	= 	"club"
	diamonds 		cardSuit 	= 	"diamond"
)

// cardValue - constants for card values
type cardValue 	string
const (
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
)

// constants for the length of arrays
const (
	// numberOfSuits - length of the suits array
	numberOfSuits	int	= 4

	// numberOfValues - length of the values array
	numberOfValues	int	= 13

	// maxScore - maximum number of points in the game
	maxScore		int	= 21
)

// the card struct describes the card.
// the struct contains the values, suits and costs of the card
type card struct {
	value 	cardValue
	suit 	cardSuit
	cost 	int //TODO разделить механики
}

// the player struct describes the player
// the struct contains the player's cards and his name
type player struct {
	cards []card
	name string
}

// the score method returns the sum of the player's points
func (p *player) score() (result int) {
	for _, card := range p.cards {
		result += card.cost
	}
	return
}

// the showCards method returns the values and suits of all the player's cards as a string
func (p *player) showCards() (allCards string) {
	for _, card := range p.cards {
		if card.value != ""{
			allCards += fmt.Sprintf("%s %s, ", card.value, card.suit)
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
	suits := [numberOfSuits]cardSuit{spades, hearts, clubs, diamonds}

	// values - an array of all values of cards
	values := [numberOfValues]cardValue{two, three, four, five, six, seven, eight, nine, ten, jack, queen, king, ace}

	// costs - an array of all costs of cards
	costs := [numberOfValues]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 2, 3, 4, 11}

	// deck - a slice of a deck of cards
	deck := make([]card, numberOfValues*numberOfSuits)

	// current - a variable for the current card in the deck
	current := 0

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
	// function for randomness
	rand.Seed(time.Now().UnixNano())

	for key := range deck {
		randInt := rand.Intn(len(deck))
		deck[key], deck[randInt] = deck[randInt], deck[key]
	}
}

// func computerGame creates a computer player and gives him cards
// input parameter: deck of cards
// output parameter: computer player, playerAI - the object of the player structure TODO исправить коммент
func computerGame(deck *[]card) (playerAI player) {
	playerAI.name = "computer"

	// costs - an array of all costs of cards
	costs := [numberOfValues]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 2, 3, 4, 11}

	// matchCards - number of matching cards
	matchCards := 0

	// chance - a chance to take another card
	chance := 0

	// function for randomness
	rand.Seed(time.Now().UnixNano())

	for range *deck {
		if playerAI.score() < 17 {
			playerAI.takeCard(deck)
		} else if playerAI.score() < 20 {
			for _, value := range costs {
				if value <= maxScore - playerAI.score() {
					matchCards += 1
				}
			}
			chance = int(float64(matchCards) / float64(numberOfValues) * 100)
			randInt := 1 + rand.Intn(100)
			if randInt <= chance {
				playerAI.takeCard(deck)
			} else {
				break
			}
		} else {
			break
		}
	}
	return
}

// func playerGame creates a user player and gives him cards
// input parameters: deck of cards, name of user
// output parameter: user player, playerUser - the object of the player structure TODO комментарий
func playerGame(deck *[]card, name string) (playerUser player, err error){
	playerUser.name = name

	if len(*deck) != 0 {
		playerUser.takeCard(deck)
		playerUser.info()
	} else {
		return playerUser, nil
	}

stop:
	for range *deck {
		// answer - a variable for the player's response
		var answer string
		fmt.Println("more? (Y/N)")
		_, err := fmt.Scanln(&answer)
		if err != nil {
			fmt.Println(err)
			return player{}, err
		}
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
	return playerUser, nil
}

// func result outputs the result of the round to the console
// input parameters: computer player, user player
func result(playerAI player, playerUser player) {
	fmt.Println("------------------------------------")
	fmt.Println("---------------RESULT---------------")
	fmt.Println("------------------------------------")

	playerAI.info()
	playerUser.info()

	if playerAI.score() > maxScore && playerUser.score() > maxScore {
		fmt.Println("everyone lost!")
	} else if (playerAI.score() > playerUser.score() && playerAI.score() <= maxScore) || (playerUser.score() > maxScore) {
		fmt.Println(playerAI.name + " won!")
	} else if (playerAI.score() < playerUser.score() && playerUser.score() <= maxScore) || (playerAI.score() > maxScore) {
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
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
		return
	}

restart:
	for range deck {
		playerAI = computerGame(&deck)
		playerUser, err = playerGame(&deck, name)
		if err != nil {
			return
		}

		result(playerAI, playerUser)

		if len(deck) == 0 {
			fmt.Println("the deck is over!")
			break
		}

		for {
			// answer - a variable for the player's response
			var answer string
			fmt.Println("start the game again? (Y/N)")
			_, err := fmt.Scanln(&answer)
			if err != nil {
				fmt.Println(err)
				return
			}
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


