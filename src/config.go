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

var config = "./config.yaml"
var env = "./.env"

type Config struct {
	Token  string `yaml:"token"`
	BotUrl string `yaml:"bot_url"`
	Admin1 string `yaml:"admin_1"`
	Admin2 string `yaml:"admin_2"`
	DB     struct {
		Url string `yaml:"db_url"`
		Key string `yaml:"db_key"`
	} `yaml:"db"`
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
		config.Token = token
	}

	if botUrl := os.Getenv("BOT_URL"); botUrl != "" {
		config.BotUrl = botUrl
	}

	if admin1 := os.Getenv("ADMIN_1"); admin1 != "" {
		config.Admin1 = admin1
	}
	if admin2 := os.Getenv("ADMIN_2"); admin2 != "" {
		config.Admin2 = admin2
	}

	if url := os.Getenv("DB_URL"); url != "" {
		config.DB.Url = url
	}

	if key := os.Getenv("DB_TOKEN"); key != "" {
		config.DB.Key = key
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
