package main

import (
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	token := os.Getenv("TOKEN")
	client, err := tb.NewBot(tb.Settings{
		URL:    "https://api.telegram.org",
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	bot := Bot{client}
	bot.client.Handle("/help", bot.helpHandler)
	bot.client.Handle("/balance", bot.balanceHandler)
	bot.client.Start()
}
