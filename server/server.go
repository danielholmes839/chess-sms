package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/twilio/twilio-go"
)

// State stored per number
type Game struct {
}

// Puzzle information
type Puzzle struct {
	Id       int    // The puzzle id, matches image name
	Solution string // The move required
	Hint     string // The name of the piece
}

type server struct {
	games   map[string]*Game
	puzzles []*Puzzle
	config  *config
	twilio  *twilio.RestClient // The twilio client
}

type config struct {
	host   string
	sender string // The phone number to send texts from
}

func (s *server) handleTwilio() http.HandlerFunc {
	// Handle twilio
	return func(w http.ResponseWriter, r *http.Request) {
		message, _ := GetTwilioMessage(r)
		command := strings.ToLower(message.Body)

		switch {
		case strings.HasPrefix(command, "commands"):
			// Help command
			s.handleInfoText(message)
			break

		case strings.HasPrefix(command, "puzzle"):
			// Puzzle command
			s.handlePuzzleText(message)
			break

		case strings.HasPrefix(command, "move"):
			// Move command
			s.handleMoveText(message)
			break

		case strings.HasPrefix(command, "hint"):
			// Hint command
			s.handleMoveText(message)
			break

		case strings.HasPrefix(command, "answer"):
			s.handleMoveText(message)
			break
		default:
			s.handleInvalidText()
		}
	}
}

func (s *server) handleInfoText(message *IncomingTwilioMessage) {
	// Send info to the user
	SendTwilioMessage(s, &TwilioMessage{
		To:   message.From,
		Body: "\n\nHack the North 2021: Chess Puzzles through Twilio!\n\nCommands:\n- commands\n- puzzle\n- move <move>\n- hint\n- answer\n",
	})
}

func (s *server) handlePuzzleText(message *IncomingTwilioMessage) {
	
}

func (s *server) handleMoveText(message *IncomingTwilioMessage) {

}

func (s *server) handleHintText(message *IncomingTwilioMessage) {

}

func (s *server) handleAnswerText(message *IncomingTwilioMessage) {

}

func (s *server) handleInvalidText() {

}

func (s *server) handleImage() http.HandlerFunc {
	// Setup handle func to serve images
	fs := http.FileServer(http.Dir("./image/data"))
	handler := http.StripPrefix("/image/", fs)
	return handler.ServeHTTP
}

func Start() {
	s := &server{
		games:   make(map[string]*Game),
		puzzles: make([]*Puzzle, 0),
		config: &config{
			host:   os.Getenv("HOST"),
			sender: os.Getenv("TWILIO_SENDER"),
		},
		twilio: twilio.NewRestClientWithParams(twilio.RestClientParams{
			Username:   os.Getenv("TWILIO_USERNAME"),
			Password:   os.Getenv("TWILIO_PASSWORD"),
			AccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		}),
	}

	// Setup handlers
	http.HandleFunc("/twilio", s.handleTwilio())
	http.HandleFunc("/image/", s.handleImage())

	fmt.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
