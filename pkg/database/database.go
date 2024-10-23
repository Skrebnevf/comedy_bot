package database

import (
	"encoding/json"
	"log"

	"github.com/supabase-community/supabase-go"
	"gopkg.in/telebot.v4"
)

type Events struct {
	Description string `json:"description"`
}

func GetEvents(db *supabase.Client) string {
	res, _, err := db.From("events").
		Select("description", "", false).
		Eq("id", "1").
		Execute()

	if err != nil {
		log.Printf("Error inserting into Supabase: %v", err)
	}

	var events []Events
	if err := json.Unmarshal(res, &events); err != nil {
		log.Printf("unmarshal response error: %v", err)
	}

	event := events[0].Description
	return event
}

func AddEvent(c telebot.Context, db *supabase.Client, desc string) {
	desc = c.Message().Text

	var result []map[string]interface{}
	if len(result) == 0 {
		insert := Events{
			Description: desc,
		}

		_, _, err := db.From("events").
			Update(insert, "representation", "").
			Eq("id", "1").
			Execute()

		if err != nil {
			log.Printf("Error inserting into Supabase: %v", err)
		}
	}
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	IsBot    bool   `json:"isBot"`
}

func WriteUser(c telebot.Context, db *supabase.Client) {
	userID := c.Sender().ID
	username := c.Sender().Username
	name := c.Sender().FirstName
	surname := c.Sender().LastName
	isBot := c.Sender().IsBot

	var result []map[string]interface{}
	if len(result) == 0 {
		insertData := User{
			ID:       userID,
			Username: username,
			Name:     name,
			Surname:  surname,
			IsBot:    isBot,
		}

		_, _, err := db.From("users").
			Insert(insertData, true, "id", "representation", "").
			Execute()
		if err != nil {
			log.Printf("Error inserting into Supabase: %v", err)
		}
	}
}
