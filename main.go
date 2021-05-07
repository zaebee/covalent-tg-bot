package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/zaebee/govalent"
	"github.com/zaebee/govalent/class_a"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	token := os.Getenv("TOKEN")
	b, err := tb.NewBot(tb.Settings{
		URL:    "https://api.telegram.org",
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/help", func(m *tb.Message) {
		msg := `Available commands:
		/help - show this message
		/balance <chain_id> <address> - show wallet balances`
		b.Send(m.Sender, msg)
	})

	b.Handle("/balance", func(m *tb.Message) {
		var message string
		params := strings.Split(m.Payload, " ")
		log.Printf("got params: %v", params)
		if len(params) < 2 {
			message = "Not enough params. Please, set chain_id and address."
		}
		if len(params) == 2 {
			var balances []string
			info, err := govalent.ClassA().GetTokenBalances(params[0], params[1], class_a.BalanceParams{Nft: false})
			if err != nil {
				message = fmt.Sprintf("Sorry, I can't get balance for given address: %v", err)
			}
			for _, i := range info.Items {
				balance := fmt.Sprintf("%v - %v %v", i.ContractName, i.Balance, i.ContractTickerSymbol)
				balances = append(balances, balance)
			}
			message = strings.Join(balances, "\n")
			message = fmt.Sprintf("Balance: %+v\n%s", info.Address, message)
		}
		b.Send(m.Sender, message)
	})

	b.Start()
}
