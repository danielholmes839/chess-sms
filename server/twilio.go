package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Incoming twilio message
type TwilioMessage struct {
	From  string
	Body  string
	Other url.Values
}

// Get incoming Twilio message
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

// Reply to a message
func (message *TwilioMessage) Reply(s *server, body string) {
	s.twilio.ApiV2010.CreateMessage(&openapi.CreateMessageParams{
		To:   &message.From,
		Body: &body,
		From: &s.config.sender,
	})
}

// Reply to a message and include a puzzle
func (message *TwilioMessage) ReplyWithPuzzle(s *server, body string, puzzleId int) {
	s.twilio.ApiV2010.CreateMessage(&openapi.CreateMessageParams{
		To:       &message.From,
		Body:     &body,
		From:     &s.config.sender,
		MediaUrl: &[]string{fmt.Sprintf("%s/image/%d.png", s.config.host, puzzleId)},
	})
}
