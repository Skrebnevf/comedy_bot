package database

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/supabase-community/supabase-go"
	"gopkg.in/telebot.v4"
)

type Events struct {
	Description string `json:"description"`
}

func AddEvent(c telebot.Context, db *supabase.Client, desc string) error {
	desc = c.Message().Text

	insert := Events{
		Description: desc,
	}

	_, _, err := db.From("events").
		Update(insert, "representation", "exact").
		Eq("id", "1").
		Execute()
	if err != nil {
		return fmt.Errorf("Error inserting into Supabase: %v", err)
	}
	return nil
}

func GetEvents(db *supabase.Client) (string, error) {
	res, _, err := db.From("events").
		Select("description", "exact", false).
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
	Reservation string `json:"reservations"`
}

func GetReservations(db *supabase.Client) (string, error) {
	res, _, err := db.From("reservations").
		Select("reservations", "exact", false).
		Eq("id", "1").
		Execute()
	if err != nil {
		return "", fmt.Errorf("Error inserting into Supabase: %v", err)
	}

	var r []Reservations
	if err := json.Unmarshal(res, &r); err != nil {
		return "", fmt.Errorf("unmarshal response error: %v", err)
	}

	reservations := r[0].Reservation
	return reservations, nil
}

func AddReservations(c telebot.Context, db *supabase.Client, msg string) error {
	var result []map[string]interface{}
	if len(result) == 0 {
		insert := Reservations{
			Reservation: msg,
		}

		_, _, err := db.From("reservations").
			Update(insert, "representation", "exact").
			Eq("id", "1").
			Execute()
		if err != nil {
			return fmt.Errorf("Error inserting into Supabase: %v", err)
		}
	}
	return nil
}

type Cancelations struct {
	Cancelation string `json:"cancelation"`
}

func GetCancelReservations(db *supabase.Client) (string, error) {
	res, _, err := db.From("cancelation_reservation").
		Select("cancelation", "exact", false).
		Eq("id", "1").
		Execute()
	if err != nil {
		return "", fmt.Errorf("Error inserting into Supabase: %v", err)
	}

	var c []Cancelations
	if err := json.Unmarshal(res, &c); err != nil {
		return "", fmt.Errorf("unmarshal response error: %v", err)
	}

	reservations := c[0].Cancelation
	return reservations, nil
}

func CancelReservation(c telebot.Context, db *supabase.Client, msg string) error {
	var result []map[string]interface{}
	if len(result) == 0 {
		insert := Cancelations{
			Cancelation: msg,
		}

		_, _, err := db.From("cancelation_reservation").
			Update(insert, "representation", "exact").
			Eq("id", "1").
			Execute()
		if err != nil {
			return fmt.Errorf("Error inserting into Supabase: %v", err)
		}
	}
	return nil
}

type MessageLog struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func WriteMessageLog(c telebot.Context, db *supabase.Client) error {
	username := c.Sender().Username
	message := c.Message().Text

	insertData := MessageLog{
		Username: username,
		Message:  message,
	}

	_, _, err := db.From("message_log").
		Insert(insertData, true, "uuid", "representation", "exact").
		Execute()
	if err != nil {
		return fmt.Errorf("Error inserting into Supabase: %v", err)
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

func GetUserIDs(db *supabase.Client) ([]User, error) {
	resp, _, err := db.From("users").
		Select("*", "exact", false).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("cannot get users from db, err: %v", err)
	}

	var u []User
	err = json.Unmarshal(resp, &u)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal user data, err: %v", err)
	}

	return u, nil
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
			Insert(insertData, true, "id", "representation", "exact").
			Execute()

		if err != nil {
			return fmt.Errorf("Error inserting into Supabase: %v", err)
		}
	}
	return nil
}

type UserCommand struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Command  string `json:"command"`
	Count    int64  `json:"count"`
}

func AddCommandCounter(c telebot.Context, db *supabase.Client) error {
	userID := c.Sender().ID
	command := c.Message().Text
	username := c.Sender().Username

	res, _, err := db.From("command_counter").
		Select("count", "exact", false).
		Eq("id", strconv.Itoa(int(userID))).
		Eq("command", command).
		Execute()
	if err != nil {
		return fmt.Errorf("cannot get command counter table, error: %v", err)
	}

	var r []UserCommand
	if err := json.Unmarshal(res, &r); err != nil {
		return fmt.Errorf("JSON-encoded error: %v", err)
	}

	if len(r) == 0 {
		insert := UserCommand{
			ID:       userID,
			Count:    1,
			Command:  command,
			Username: username,
		}

		_, _, err := db.From("command_counter").
			Insert(insert, false, "uuid", "minimal", "exact").
			Execute()
		if err != nil {
			return fmt.Errorf("cannot create new record in command_counter table, error: %v", err)
		}
	} else {
		newCount := r[0].Count + 1
		insert := UserCommand{
			ID:       userID,
			Count:    newCount,
			Command:  command,
			Username: username,
		}

		_, _, err := db.From("command_counter").
			Update(insert, "minimal", "exact").
			Eq("id", strconv.Itoa(int(userID))).
			Eq("command", command).
			Execute()
		if err != nil {
			return fmt.Errorf("cannot set new counter value, error: %v", err)
		}
	}
	return err
}
