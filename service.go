package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func sendMessageToTelegram(chatID string, message string) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	formData := url.Values{
		"chat_id":    {chatID},
		"text":       {message},
		"parse_mode": {"Markdown"},
	}
	resp, err := http.PostForm(endpoint, formData)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	if result["ok"].(bool) {
		return nil
	} else {
		return errors.New(result["description"].(string))
	}
}

func setTelegramWebhookPath(path string) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", botToken, url.QueryEscape(path))
	resp, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	if result["ok"].(bool) {
		return nil
	} else {
		return errors.New(result["description"].(string))
	}
}
