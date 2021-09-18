package main

import (
	"fmt"

	"./game"
)

func main() {
	fmt.Println("The game has started!")

	deck := game.CreateDeck()
	game.RandomDeck(deck)
	game.Game(deck)

	fmt.Println("The game is over!")
}


