package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/notnil/chess"
)

type Puzzle struct {
	id      int // The puzzle id, matches image name
	answer  string
	answers []string // The correct move(s) in different formats
	hint    string   // The name of the piece
	color   string   // The color to move "White" or "Black"
	game    *chess.Game
}

func (p *Puzzle) IsCorrect(move string) bool {
	move = strings.ToLower(move)
	for _, option := range p.answers {
		if move == option {
			return true
		}
	}
	return false
}

func (p *Puzzle) GetID() int {
	return p.id
}

func (p *Puzzle) GetHint() string {
	return p.hint
}

func (p *Puzzle) GetAnswer() string {
	return p.answer
}

func (p *Puzzle) GetDescription() string {
	return fmt.Sprintf("Find the checkmate in one. %s to move", p.color)
}

func (p *Puzzle) GetGame() *chess.Game {
	return p.game
}

func GetPuzzles() []*Puzzle {
	return []*Puzzle{{
		id:      1,
		answer:  "d5",
		answers: []string{"d5", "d7d5"},
		hint:    "Pawn",
		color:   "Black",
	}}
}

type PuzzleJSON struct {
	Answer string `json:"answer"`
	Color  string `json:"color"`
	PGN    string `json:"pgn"`
}

func (p PuzzleJSON) ConvertToPuzzle(id int) *Puzzle {
	games, _ := chess.GamesFromPGN(strings.NewReader(p.PGN))
	game := games[0]

	base := strings.ToLower(p.Answer)
	base2 := strings.Replace(base, "#", "", 1)

	return &Puzzle{
		id:      id,
		answer:  p.Answer,
		answers: []string{base, base2},
		hint:    "IMPLEMENT",
		color:   p.Color,
		game:    game,
	}
}

func ReadPuzzles() []*Puzzle {
	f, _ := os.Open("python/puzzles.json")
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
