package main

import (
	"fmt"
	"htn/server/game"
	"htn/server/image"
)

func main() {
	puzzles := game.ReadPuzzles()

	for _, puzzle := range puzzles {
		image.Generate(puzzle)
		fmt.Println(puzzle.GetID(), puzzle.GetAnswer(), puzzle.GetDescription(), puzzle.GetHint())
	}

}
