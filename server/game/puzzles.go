package game

import (
	"fmt"
	"strings"
)

type Puzzle struct {
	id      int // The puzzle id, matches image name
	answer  string
	answers []string // The correct move(s) in different formats
	hint    string   // The name of the piece
	move    string   // The color to move "White" or "Black"
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
	return fmt.Sprintf("Find the checkmate in one. %s to move", p.move)
}

func GetPuzzles() []*Puzzle {
	return []*Puzzle{{
		id:      1,
		answer:  "e5",
		answers: []string{"d5", "d7d5"},
		hint:    "Pawn",
		move:    "Black",
	}}
}
