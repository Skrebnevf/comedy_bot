package handlers

import (
	"io"
	"log"
	"net/http"
	"time"
)

func WakeUp(url string) {
	for {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Ошибка при запросе: %v", err)
		} else {
			log.Printf("Запрос выполнен, статус: %s", resp.Status)
			resp.Body.Close()
		}
		time.Sleep(1 * time.Second)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("cannot get body request: %v", err)
		}

		log.Printf("Resp: %v, Time %v", body, time.Now())
	}
}
