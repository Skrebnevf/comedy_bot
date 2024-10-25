package database

import (
	"encoding/json"
	"fmt"

	"github.com/supabase-community/supabase-go"
	"gopkg.in/telebot.v4"
)

type Events struct {
	Description string `json:"description"`
}

func AddEvent(c telebot.Context, db *supabase.Client, desc string) error {
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
			return fmt.Errorf("Error inserting into Supabase: %v", err)
		}
	}

	return nil
}

func GetEvents(db *supabase.Client) (string, error) {
	res, _, err := db.From("events").
		Select("description", "", false).
		Eq("id", "1").
		Execute()
	if err != nil {
		return "", fmt.Errorf("Error inserting into Supabase: %v", err)
	}

	var e []Events
	if err := json.Unmarshal(res, &e); err != nil {
		return "", fmt.Errorf("unmarshal response error: %v", err)
	}

	event := e[0].Description
	return event, nil
}

type Reservations struct {
	Description string `json:"reservations"`
}

func GetReservations(db *supabase.Client) (string, error) {
	res, _, err := db.From("reservations").
		Select("reservations", "", false).
		Eq("id", "1").
		Execute()
	if err != nil {
		return "", fmt.Errorf("Error inserting into Supabase: %v", err)
	}

	var r []Reservations
	if err := json.Unmarshal(res, &r); err != nil {
		return "", fmt.Errorf("unmarshal response error: %v", err)
	}

	reservations := r[0].Description
	return reservations, nil
}

func AddReservations(c telebot.Context, db *supabase.Client, msg string) error {
	var result []map[string]interface{}
	if len(result) == 0 {
		insert := Reservations{
			Description: msg,
		}

		_, _, err := db.From("reservations").
			Update(insert, "representation", "").
			Eq("id", "1").
			Execute()

		if err != nil {
			return fmt.Errorf("Error inserting into Supabase: %v", err)
		}
	}

	return nil
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	IsBot    bool   `json:"isBot"`
}

func WriteUser(c telebot.Context, db *supabase.Client) error {
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
			return fmt.Errorf("Error inserting into Supabase: %v", err)
		}
	}
	return nil
}
