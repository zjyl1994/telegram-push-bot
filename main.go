package main

import (
	"log"
	"net/http"
	"os"
)

var botToken string

func main() {
	log.Println("OhMyPushBot")
	botToken = os.Getenv("TELEGRAM_PUSH_BOT_TOKEN")
	http.HandleFunc("/telegram/webhook", telegramWebhookHandler)
	http.HandleFunc("/send", sendMessageWebhookHandler)
	if err := setTelegramWebhookPath(telegramWebhookURLGen()); err != nil {
		log.Fatalln("set telegram webhook fail", err.Error())
		return
	}
	if err := http.ListenAndServe(os.Getenv("TELEGRAM_PUSH_BOT_PORT"), nil); err != nil {
		log.Fatalln("server run fail", err.Error())
	} else {
		log.Println("server stoped.")
	}
}
