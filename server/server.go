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
	users   *game.UserManager
	puzzles []*game.Puzzle
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
		user := s.users.Get(message.From)

		switch {
		case strings.HasPrefix(command, "commands"):
			// Help/info command
			info := "\n\nHack the North 2021: Chess Puzzles through Twilio!\n\nCommands:\n- commands\n- puzzle\n- move <move>\n- hint\n- answer\n"
			message.Reply(s, info)
			break

		case strings.HasPrefix(command, "puzzle"):
			// Puzzle command
			s.handlePuzzleText(user, message)
			break

		case strings.HasPrefix(command, "move"):
			// Move command
			s.handleMoveText(user, message)
			break

		case strings.HasPrefix(command, "hint"):
			// Hint command
			s.handleHintText(user, message)
			break

		case strings.HasPrefix(command, "answer"):
			// Answer command
			s.handleAnswerText(user, message)
			break

		default:
			// Suggest sending 'commands'
			message.Reply(s, "Sorry that command doesn't exist. Try sending 'commands' for more information")
		}
	}
}

func (s *server) handleImage() http.HandlerFunc {
	// Setup handle func to serve images
	fs := http.FileServer(http.Dir("./image/data"))
	handler := http.StripPrefix("/image/", fs)
	return handler.ServeHTTP
}

func Start() {
	s := &server{
		users:   game.NewUserManager(),
		puzzles: game.GetPuzzles(),
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