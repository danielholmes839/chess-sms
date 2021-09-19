package server

import (
	"errors"
	"fmt"
	"htn/server/game"
	"math/rand"
	"strings"
)

func (s *server) handlePuzzleText(user *game.User, message *IncomingTwilioMessage) {
	if user != nil {
		message.Reply(s, "Sorry you've already started a puzzle. Try sending 'answer' before starting a new puzzle")
		return
	}

	// Create the new user with a puzzle for them to solve
	puzzle := s.puzzles[rand.Intn(500)]
	user = game.NewUser(puzzle)

	s.users.Add(message.From, user)
	message.ReplyWithPuzzle(s, puzzle.GetDescription(), puzzle.GetID())

}

// Parses the move command to extract the intended move
func parseMove(command string) (string, error) {
	command = strings.ToLower(command)
	args := strings.SplitAfter(command, "move")

	if len(args) < 2 {
		return "", errors.New("Please enter move. For example: 'move Nf3' or 'move g1f3'")
	}

	return strings.TrimSpace(args[1]), nil
}

func (s *server) handleMoveText(user *game.User, message *IncomingTwilioMessage) {
	// Check if the move entered by the user was correct
	if user == nil {
		message.Reply(s, "Sorry you need to start a puzzle before making a move. Try sending 'puzzle'")
		return
	}

	// Check the move is correct to continue
	puzzle := user.GetPuzzle()
	if !puzzle.IsCorrect(message.Body) {
		incorrect := fmt.Sprintf("'%s' is incorrect try again!", message.Body)
		message.Reply(s, incorrect)
		return
	}

	// Reply with correct!
	correct := fmt.Sprintf("'%s' is correct!", puzzle.GetAnswer())
	message.Reply(s, correct)
	s.users.Remove(message.From) // disconnect
}

func (s *server) handleHintText(user *game.User, message *IncomingTwilioMessage) {
	// Send the user a hint
	if user == nil {
		message.Reply(s, "Sorry you need to start a puzzle before getting a hint. Try sending 'puzzle'")
		return
	}

	message.Reply(s, user.GetPuzzle().GetHint())
}

func (s *server) handleAnswerText(user *game.User, message *IncomingTwilioMessage) {
	/// Send the user the answer, and disconnect
	if user == nil {
		message.Reply(s, "Sorry you need to start a puzzle before getting an answer. Try sending 'puzzle'")
		return
	}

	message.Reply(s, user.GetPuzzle().GetAnswer())
	s.users.Remove(message.From) // disconnect
}
