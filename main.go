package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	dotenv "github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func ClientExample() {
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username:   "SK56b6739ee32e4247b20db49f44fe935b",
		Password:   "PKW5YCDaG9nMNoEYog138IX0yCRxILsU",
		AccountSid: "AC41ffe205569ce9fd21a8bc6d22546452",
	})

	to := "613-299-1379"
	body := "Solve this chess position!"
	from := "613-704-4683"
	img := []string{"https://api.stop-checker.com/image/1.png"}

	client.ApiV2010.CreateMessage(&openapi.CreateMessageParams{
		To:       &to,
		Body:     &body,
		From:     &from,
		MediaUrl: &img,
	})
}

type TwilioMessage struct {
	From  string
	Body  string
	Other url.Values
}

func GetTwilioMessage(r *http.Request) (*TwilioMessage, error) {
	buf := new(strings.Builder)
	io.Copy(buf, r.Body)
	body := buf.String()

	params, err := url.ParseQuery(body)
	if err != nil {
		return nil, err
	}

	return &TwilioMessage{
		From:  params.Get("From"),
		Body:  params.Get("Body"),
		Other: params,
	}, nil
}

func response(w http.ResponseWriter, message string) {
	w.Write([]byte(message))
}

func TwilioHandler(w http.ResponseWriter, r *http.Request) {
	message, err := GetTwilioMessage(r)

	if err != nil {
		response(w, "Error could not parse text message")
		return
	}

	if message.Body == "puzzle" {
		ClientExample()
	}

	fmt.Printf("%s: %s\n", message.From, message.Body)
}

// State stored per number
type Game struct {
}

// Puzzle information
type Puzzle struct {
	Id       string // The puzzle id, matches image name
	Solution string // The move required
	Hint     string // The name of the piece
}

type server struct {
	games   map[string]*Game
	puzzles []*Puzzle
	twilio  *twilio.RestClient
}

func (s *server) handleTwilio() http.HandlerFunc {
	// Handle twilio
	return func(w http.ResponseWriter, r *http.Request) {
		message, err := GetTwilioMessage(r)

		if err != nil {
			response(w, "Error could not parse text message")
			return
		}

		if message.Body == "puzzle" {
			ClientExample()
		}

		response(w, "Echo: "+message.Body)
	}
}

func (s *server) handleImage() http.HandlerFunc {
	// Setup handle func to serve images
	fs := http.FileServer(http.Dir("./image/data"))
	handler := http.StripPrefix("/image/", fs)
	return handler.ServeHTTP
}

func main() {
	// Read .env
	dotenv.Load(".env")

	// Server config
	s := &server{
		games:   make(map[string]*Game),
		puzzles: make([]*Puzzle, 0),
		twilio: twilio.NewRestClientWithParams(twilio.RestClientParams{
			Username:   os.Getenv("TWILIO_USERNAME"),
			Password:   os.Getenv("TWILIO_PASSWORD"),
			AccountSid: os.Getenv("TIWLIO_ACCOUNT_SID"),
		}),
	}

	// Setup handlers
	http.HandleFunc("/twilio", s.handleTwilio())
	http.HandleFunc("/image/", s.handleImage())
	http.ListenAndServe(":3000", nil)
}
