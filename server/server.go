package server

import (
	"fmt"
	"htn/server/game"
	"net/http"
	"os"
	"strings"

	"github.com/twilio/twilio-go"
)

type server struct {
	users   *game.PuzzleManager
	puzzles []*game.Puzzle
	config  *config
	twilio  *twilio.RestClient // The twilio client
}

type config struct {
	host   string
	sender string // The phone number to send texts from
}

func (s *server) handleTwilio() http.HandlerFunc {
	// Handle twilio SMS
	return func(w http.ResponseWriter, r *http.Request) {
		message, _ := GetTwilioMessage(r)
		command := strings.ToLower(message.Body)
		user := s.users.Get(message.From)

		switch {
		case strings.HasPrefix(command, "commands"):
			// Help/info command
			s.replyHelp(message)
			break

		case strings.HasPrefix(command, "puzzle"):
			// Puzzle command
			s.handlePuzzleText(user, message)
			break

		case strings.HasPrefix(command, "hint"):
			// Hint command
			s.handleHintText(user, message)
			break

		case strings.HasPrefix(command, "answer"):
			// Answer command
			s.handleAnswerText(user, message)
			break

		case user != nil:
			// The user sent a move
			s.handleMoveText(user, message)
			break

		default:
			// Suggest sending 'commands'
			s.replyHelpPrompt(message)
		}
	}
}

func (s *server) handleImage() http.HandlerFunc {
	// Serve images
	fs := http.FileServer(http.Dir("./server/image/data/"))
	handler := http.StripPrefix("/image/", fs)
	return handler.ServeHTTP
}

func Setup() {
	s := &server{
		users:   game.NewPuzzleManager(),
		puzzles: game.ReadPuzzles(),
		config: &config{
			host:   os.Getenv("HOST"),
			sender: os.Getenv("TWILIO_SENDER"),
		},
		twilio: twilio.NewRestClientWithParams(twilio.RestClientParams{
			Username: os.Getenv("TWILIO_SID"),
			Password: os.Getenv("TWILIO_PASSWORD"),
		}),
	}

	// Setup handlers
	http.HandleFunc("/twilio", s.handleTwilio())
	http.HandleFunc("/image/", s.handleImage())

	fmt.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
