package handlers

import (
	"io"
	"log"
	"net/http"
)

func WakeUp(url string, command string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("bot is not wake up: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("cannot read body from wake up: %v", err)
	}

	log.Println(string(body) + command)
}
