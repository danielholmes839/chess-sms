package game

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/notnil/chess"
)

type PuzzleJSON struct {
	Answer string `json:"answer"`
	Color  string `json:"color"`
	PGN    string `json:"pgn"`
	Hint   string `json:"hint"`
}

func (p PuzzleJSON) ConvertToPuzzle(id int) *Puzzle {
	games, _ := chess.GamesFromPGN(strings.NewReader(p.PGN))
	game := games[0]

	positions := game.Positions()
	position := positions[len(positions)-2]

	moves := game.Moves()
	move := moves[len(moves)-1]

	pgn1 := strings.ToLower(p.Answer)
	pgn2 := strings.Replace(pgn1, "#", "", 1)
	s1s2 := move.String()

	return &Puzzle{
		id:       id,
		answer:   p.Answer,
		answers:  []string{pgn1, pgn2, s1s2},
		hint:     p.Hint,
		color:    p.Color,
		position: position,
	}
}

func ReadPuzzles() []*Puzzle {
	f, _ := os.Open("dataset/puzzles.json")
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	puzzleJSON := []PuzzleJSON{}
	json.Unmarshal(data, &puzzleJSON)

	puzzles := make([]*Puzzle, len(puzzleJSON))
	for i, puzzle := range puzzleJSON {
		puzzles[i] = puzzle.ConvertToPuzzle(i + 1)
	}
	return puzzles
}
