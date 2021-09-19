package main

import (
	"fmt"
	"htn/server/game"
)

func main() {
	puzzles := game.ReadPuzzles()
	fmt.Println(puzzles[0].GetDescription(), puzzles[0].GetGame().Position().Board().Draw())

	for _, puzzle := range puzzles {
		game := puzzle.GetGame()
		game.Positions()
		fmt.Println(game.Moves())
	}

}
