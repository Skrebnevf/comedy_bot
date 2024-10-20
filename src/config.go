package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
	"gopkg.in/yaml.v3"
)

var config = "config.yaml"
var env = ".env"

type Config struct {
	Tokens struct {
		TestToken   string `yaml:"test"`
		ComedyToken string `yaml:"comedy"`
	} `yaml:"token"`
}

func LoadConfig() (*Config, error) {
	file, err := os.ReadFile(config)
	if err != nil {
		return nil, fmt.Errorf("cannot read yaml file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	if err := godotenv.Load(env); err != nil {
		log.Printf("Warning: could not load .env file, using default environment variables")
	}

	if token := os.Getenv("TOKEN"); token != "" {
		config.Tokens.TestToken = token
	}

	return &config, nil
}

func BotInit(token string) (*telebot.Bot, error) {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	return telebot.NewBot(pref)
}
