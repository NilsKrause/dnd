package bot

import (
	"log"
	"os"
	"os/signal"

	"de.nilskrau.dndbot/app"
	"de.nilskrau.dndbot/command"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	app      *app.App
	session  *discordgo.Session
	commands map[string]command.Command
}

func NewBot(app *app.App, commands []command.Command, appid string, guildid string, token string) *Bot {
	cmds := make([]*discordgo.ApplicationCommand, 0)
	cmdMap := make(map[string]command.Command)
	for _, cmd := range commands {
		cmds = append(cmds, cmd.GetDefinition())
		cmdMap[cmd.GetName()] = cmd
	}

	bot := &Bot{
		app:      app,
		commands: cmdMap,
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	bot.session = session
	bot.session.AddHandler(bot.rootHandler)

	_, err = bot.session.ApplicationCommandBulkOverwrite(appid, guildid, cmds)
	if err != nil {
		log.Fatalf("could not register commands: %s", err)
	}
	return bot
}

func (b *Bot) GetApp() *app.App {
	return b.app
}

func (b *Bot) GetSession() *discordgo.Session {
	return b.session
}

func (b *Bot) Start() {
	err := b.session.Open()
	if err != nil {
		log.Fatalf("could not open session: %s", err)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	err = b.session.Close()
	if err != nil {
		log.Printf("could not close session gracefully: %s", err)
	}
}

func (b *Bot) rootHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()

	command, ok := b.commands[data.Name]
	if !ok || command == nil {
		panic("command not found")
	}

	handler := command.GetHandler(b, i)

	if handler == nil {
		panic("command did not return handler")
	}

	handler()
}
