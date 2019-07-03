package main

import (
	"log"
	"os"
)

var botToken string

func main() {
	log.Println("OhMyPushBot")
	botToken = os.Getenv("TELEGRAM_PUSH_BOT_TOKEN")
	err := sendMessageToTelegram("0", `*DON'T PANIC!*
This is test message
for format _check_
	`)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
