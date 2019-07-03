package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func telegramWebhookHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	chatid, text, err := parseTelegramWebhook(b)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	if strings.HasPrefix(text, "/") { // check is command
		command := strings.TrimPrefix(text, "/")
		switch command {
		case "ping":
			sendMessageToTelegram(chatid, "pong!")
		case "chatid":
			sendMessageToTelegram(chatid, chatid)
		case "url":
			sendMessageToTelegram(chatid, sendMessageURLGen(chatid))
		case "start":
			message := "*OhMyPushBot*\nCommand:\n/ping: check bot alive\n/chatid: get chat_id for this chat\n/url: get url for send message"
			sendMessageToTelegram(chatid, message)
		default:
			sendMessageToTelegram(chatid, "command not found")
		}
	}
}

func sendMessageWebhookHandler(w http.ResponseWriter, r *http.Request) {
	chatid := r.URL.Query().Get("chatid")
	sign := r.URL.Query().Get("sign")
	if len(chatid) == 0 || len(sign) == 0 {
		http.Error(w, "missing chatid or sign", 400)
		return
	}
	if !signedStringCheck(chatid, sign, botToken) {
		http.Error(w, "sign mismatch", 401)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = sendMessageToTelegram(chatid, string(b))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, "ok")
}
