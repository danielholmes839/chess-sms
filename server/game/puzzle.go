package game

import (
	"fmt"
	"strings"

	"github.com/notnil/chess"
)

type Puzzle struct {
	id       int      // The puzzle id, matches image name
	answer   string   // The answer displayed to the user
	answers  []string // The correct move(s) in different formats "Qxg7#", "Qxg7", "g1g7"
	hint     string   // The name of the piece
	color    string   // The color to move "White" or "Black"
	position *chess.Position
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

func (p *Puzzle) GetPosition() *chess.Position {
	return p.position
}
