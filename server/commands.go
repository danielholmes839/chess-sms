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

func (s *server) handlePuzzleText(puzzle *game.Puzzle, message *TwilioMessage) {
	if puzzle != nil {
		message.Reply(s, PUZZLE_IN_PROGRESSS)
		return
	}

	// Create the new user with a puzzle for them to solve
	puzzle = s.puzzles[rand.Intn(500)]

	s.users.Add(message.From, puzzle)
	message.ReplyWithPuzzle(s, puzzle.GetDescription(), puzzle.GetID())

}

func (s *server) handleMoveText(puzzle *game.Puzzle, message *TwilioMessage) {
	// Check if the move entered by the user was correct
	if puzzle == nil {
		message.Reply(s, PUZZLE_NOT_IN_PROGRESS)
		return
	}

	// Check the move is correct to continue
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

func (s *server) handleHintText(puzzle *game.Puzzle, message *TwilioMessage) {
	// Send the user a hint
	if puzzle == nil {
		message.Reply(s, PUZZLE_NOT_IN_PROGRESS)
		return
	}

	message.Reply(s, puzzle.GetHint())
}

func (s *server) handleAnswerText(puzzle *game.Puzzle, message *TwilioMessage) {
	/// Send the user the answer, and remove them
	if puzzle == nil {
		message.Reply(s, PUZZLE_NOT_IN_PROGRESS)
		return
	}

	message.Reply(s, puzzle.GetAnswer())
	s.users.Remove(message.From)
}
