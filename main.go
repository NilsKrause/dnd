package main

import (
	"flag"
	"log"
	"os"

	"de.nilskrau.dndbot/app"
	"de.nilskrau.dndbot/bot"
	"de.nilskrau.dndbot/command"
	"de.nilskrau.dndbot/command/character"
	"de.nilskrau.dndbot/command/write"
	"github.com/joho/godotenv"
)

var (
	Token = flag.String("token", "MTQyODMwODc1OTQ3MzY4ODU5Ng.GJZsv5.dV-hc53jNgc7mlh2pWh9g4tgYhJYiYeCQxiOzI", "Bot authentication token")
	App   = flag.String("app", "1428308759473688596", "Application ID")
	//Guild = flag.String("guild", "699194969226870805", "Guild ID")
	Guild = flag.String("guild", "", "Guild ID")
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  Token := os.Getenv("TOKEN")
  App := os.Getenv("APP")
  Guild := os.Getenv("GUILD")

	commands := []command.Command{
		character.CreateCharacterCommand,
		character.ListCharacterCommand,
		character.DeleteCharacterCommand,
		character.EditCharacterCommand,
		character.SetActiveCharacterCommand,
		write.WriteIncharacterCommand,
	}
	b := bot.NewBot(app.NewApp(), commands, App, Guild, Token)

	b.Start()
}
