package main

import (
	"fmt"
	"github/skrebnevf/comedy_belgrade_bot/pkg/handlers"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/supabase-community/supabase-go"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	b, err := BotInit(config.Token)
	if err != nil {
		log.Fatalf("cannot init bot, error: %v", err)
	}

	handlers.Admin1 = config.Admin1
	handlers.Admin2 = config.Admin2
	handlers.InitAllowedUsers()

	client, err := supabase.NewClient(config.DB.Url, config.DB.Key, &supabase.ClientOptions{})
	if err != nil {
		log.Printf("DB error: %v", err)
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

	go func() {
		for {
			resp, err := http.Get(config.BotUrl)
			if err != nil {
				log.Printf("Request error: %v", err)
			} else {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Printf("cannot get body request: %v", err)
				} else {
					log.Printf("Resp: %s, Time %v", body, time.Now())
				}
			}
			time.Sleep(30 * time.Second)
		}
	}()

	handlers.CommandHandlers(b, client, config.BotUrl)
	handlers.TextHandler(b, client, config.BotUrl)
	handlers.OtherHandlers(b)
	handlers.ReplyHandler(b)
	b.Start()
}
