package main

import (
	"log"
	"os"

	"github.com/ErneDainesi/discord-music-bot/bot"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    token := os.Getenv("BOT_TOKEN")
    bot.Run(token)
}
