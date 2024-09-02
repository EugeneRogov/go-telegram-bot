package main

import (
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	// tgClient := telegram.New(tgBotHost, mustToken())
	// t := mustToken()
	// fetcher = telegram.New(token)

	// fmt.Println("Привет, Мир!") // ВЫВОД: Привет, Мир!
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
