package server

import (
	"fmt"
	"htn/server/game"
	"math/rand"
)

// Messages
const HELP = "\n\nHack the North 2021: Chess Puzzles through Twilio!\n\nCommands:\n- commands\n- puzzle\n- <move>\n- hint\n- answer\n"
const HELP_PROMPT = "Sorry that command doesn't exist. Try sending 'commands' for more information"
const PUZZLE_NOT_IN_PROGRESS = "Sorry you need to start a puzzle before making a move. Try sending 'puzzle'"
const PUZZLE_IN_PROGRESSS = "Sorry you've already started a puzzle. Try sending 'answer' before starting a new puzzle"

func (s *server) replyHelp(message *TwilioMessage) {
	message.Reply(s, HELP)
}

func (s *server) replyHelpPrompt(message *TwilioMessage) {
	message.Reply(s, HELP_PROMPT)
}

func (s *server) handlePuzzleText(user *game.User, message *TwilioMessage) {
	if user != nil {
		message.Reply(s, PUZZLE_IN_PROGRESSS)
		return
	}

	// Create the new user with a puzzle for them to solve
	puzzle := s.puzzles[rand.Intn(500)]
	user = game.NewUser(puzzle)

	s.users.Add(message.From, user)
	message.ReplyWithPuzzle(s, puzzle.GetDescription(), puzzle.GetID())

}

func (s *server) handleMoveText(user *game.User, message *TwilioMessage) {
	// Check if the move entered by the user was correct
	if user == nil {
		message.Reply(s, PUZZLE_NOT_IN_PROGRESS)
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

func (s *server) handleHintText(user *game.User, message *TwilioMessage) {
	// Send the user a hint
	if user == nil {
		message.Reply(s, PUZZLE_NOT_IN_PROGRESS)
		return
	}

	message.Reply(s, user.GetPuzzle().GetHint())
}

func (s *server) handleAnswerText(user *game.User, message *TwilioMessage) {
	/// Send the user the answer, and disconnect
	if user == nil {
		message.Reply(s, PUZZLE_NOT_IN_PROGRESS)
		return
	}

	message.Reply(s, user.GetPuzzle().GetAnswer())
	s.users.Remove(message.From) // disconnect
}
