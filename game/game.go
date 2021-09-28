package game

import (
	"fmt"
	"math/rand"
	"time"

	"../card"
)

// maxScore - maximum number of points in the game
const maxScore int = 21

// computerGame creates a computer player and gives him cards
func computerGame(deck *[]card.Card) (playerAI player) {
	playerAI.name = "computer"

	// costs - an array of all costs of cards
	costs := [card.NumberOfValues]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 2, 3, 4, 11}

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
				if value <= (maxScore - playerAI.score()) {
					matchCards += 1
				}
			}
			chance = int(float64(matchCards) / float64(card.NumberOfValues) * 100)
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

// playerGame creates a user player and gives him cards
func playerGame(deck *[]card.Card, userName string) (playerUser player, err error) {
	playerUser.name = userName

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
		switch answer {
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

// result outputs the result of the round to the console
func result(playerAI player, playerUser player) {
	fmt.Println("---------------RESULT---------------")

	playerAI.info()
	playerUser.info()

	if playerAI.score() > maxScore && playerUser.score() > maxScore {
		fmt.Println("everyone lost!")
	} else if (playerAI.score() > playerUser.score() && playerAI.score() <= maxScore) || (playerUser.score() > maxScore) {
		fmt.Printf("%s won!\n", playerAI.name)
	} else if (playerAI.score() < playerUser.score() && playerUser.score() <= maxScore) || (playerAI.score() > maxScore) {
		fmt.Printf("%s won!\n", playerUser.name)
	} else if playerAI.score() == playerUser.score() {
		fmt.Println("draw!")
	}
}

// Game performs the main function of the game
func Game() {
	deck := card.CreateDeck()
	card.RandomDeck(deck)

	// playerAI, playerUser - the object of the player structure
	var playerAI, playerUser player

	// name - name of user
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
