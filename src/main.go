package main

import (
	"github/skrebnevf/comedy_belgrade_bot/pkg/handlers"
	"log"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	b, err := BotInit(config.Tokens.TestToken)
	if err != nil {
		log.Fatalf("cannot init bot, error: %v", err)
	}

	handlers.CommandHandlers(b)
	handlers.TextHandler(b)
	handlers.OtherHandlers(b)
	handlers.ReplyHandler(b)

	b.Start()
}
