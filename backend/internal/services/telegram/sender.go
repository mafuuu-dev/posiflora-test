package service_telegram

import (
	"backend/core/pkg/scope"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TelegramMessage struct {
	Text   string `json:"text"`
	ChatID string `json:"chat_id"`
}

type TelegramSender struct {
	scope *scope.Scope
}

func NewTelegramSender(scope *scope.Scope) *TelegramSender {
	return &TelegramSender{
		scope: scope,
	}
}

func (s *TelegramSender) SendMessage(botToken string, chatID string, message string) error {
	body, err := json.Marshal(TelegramMessage{
		ChatID: chatID,
		Text:   message,
	})
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram returned status: %s", resp.Status)
	}

	return nil
}
