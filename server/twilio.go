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

// Send an SMS message
func SendTwilioMessage(s *server, m *TwilioMessage) error {
	var img []string
	if m.PuzzleID != 0 {
		// Add the image url if a puzzle was specified
		img = []string{fmt.Sprintf("%s/image/%x.png", s.config.host, 1)}
	}

	// Send the message
	_, err := s.twilio.ApiV2010.CreateMessage(&openapi.CreateMessageParams{
		To:       &m.To,
		Body:     &m.Body,
		From:     &s.config.sender,
		MediaUrl: &img,
	})

	return err
}
