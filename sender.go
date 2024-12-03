package notifywork

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Sender struct {
	apiKey    string
	senderId  uint
	channelId string
}

func NewSender(apiKey string, senderId uint) *Sender {
	return &Sender{
		apiKey:   apiKey,
		senderId: senderId,
	}
}

func (s *Sender) SetDefaultChannel(channelId string) {
	s.channelId = channelId
}

func (s *Sender) SendMessage(message string) error {
	payload := map[string]any{
		"sender_id":  s.senderId,
		"channel_id": s.channelId,
		"message":    message,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, apiUrl+"/api/send", bytes.NewBuffer(payloadJson))
	req.Header.Set("Authorization", s.apiKey)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		return errors.New(string(resBody))
	}
	return nil
}
