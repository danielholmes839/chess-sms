package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Twilio incoming message helper
type IncomingTwilioMessage struct {
	From  string
	Body  string
	Other url.Values
}

// Get incoming Twilio messages
func GetTwilioMessage(r *http.Request) (*IncomingTwilioMessage, error) {
	buf := new(strings.Builder)
	io.Copy(buf, r.Body)
	body := buf.String()

	params, err := url.ParseQuery(body)
	if err != nil {
		return nil, err
	}

	return &IncomingTwilioMessage{
		From:  params.Get("From"),
		Body:  params.Get("Body"),
		Other: params,
	}, nil
}

// Twilio message being sent
type TwilioMessage struct {
	To       string
	Body     string
	PuzzleID int // 0 -> no puzzle
}

func (message *IncomingTwilioMessage) Reply(s *server, body string) {
	// Reply to an incoming message
	s.twilio.ApiV2010.CreateMessage(&openapi.CreateMessageParams{
		To:   &message.From,
		Body: &body,
		From: &s.config.sender,
	})
}

func (message *IncomingTwilioMessage) ReplyWithPuzzle(s *server, body string, puzzleId int) {
	// Reply to an incoming message, and include a puzzle
	s.twilio.ApiV2010.CreateMessage(&openapi.CreateMessageParams{
		To:       &message.From,
		Body:     &body,
		From:     &s.config.sender,
		MediaUrl: &[]string{fmt.Sprintf("%s/image/%d.png", s.config.host, puzzleId)},
	})
}
