package main

import (
	"fmt"
	"github/skrebnevf/comedy_belgrade_bot/pkg/handlers"
	"log"
	"net/http"
	"os"
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

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Bot is running")
		})

		port := os.Getenv("PORT")
		if port == "" {
			port = "8000"
		}

		log.Printf("Listening on port %s for health checks...", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal(err)
		}
	}()

	handlers.CommandHandlers(b)
	handlers.TextHandler(b)
	handlers.OtherHandlers(b)
	handlers.ReplyHandler(b)

	b.Start()
}
