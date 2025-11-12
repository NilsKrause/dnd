package command

import (
	"de.nilskrau.dndbot/app"
	"de.nilskrau.dndbot/models"
	"github.com/bwmarrin/discordgo"
)

type Options = map[string]*discordgo.ApplicationCommandInteractionDataOption

type CommandHandler func(ctx Context)
type BotHandler func()

type Command interface {
	GetName() string
	GetDefinition() *discordgo.ApplicationCommand
	GetHandler(context BotContext, interactionCreate *discordgo.InteractionCreate) BotHandler
}

type BotContext interface {
	GetApp() *app.App
	GetSession() *discordgo.Session
}

type Context interface {
	BotContext
	GetOptions() Options
	GetOption(name string) *discordgo.ApplicationCommandInteractionDataOption
	GetUser() *discordgo.User
	SendMessage(text string) error
	GetPlayer() (*models.Player, error)
	GetUserCharacters() ([]*models.Character, error)
	GetUserActiveCharacter() (*models.Character, error) 
	GetUserCharacterByHandle(handle string) (*models.Character, error) 
	SendIncharacterEmbed(char *models.Character, player *models.Player, message string) error
	SendErrorMessage(message string) error 
	SendInfo(message string) error
	SendCharacterEmbeds(chars []*models.Character) error 
	GetCurrentServerId() string
	CreateCharacter(name string, handle string, picture string, active bool) error 
}

type command struct {
	handler    CommandHandler
	definition *discordgo.ApplicationCommand
}

func (c *command) GetDefinition() *discordgo.ApplicationCommand {
	return c.definition
}

func (c *command) GetHandler(ctx BotContext, i *discordgo.InteractionCreate) BotHandler {
	cc := buildCommandContext(ctx, i)
	return func() {
		c.handler(cc)
	}
}

func (c *command) GetName() string {
	return c.definition.Name
}

func NewCommand(handler CommandHandler, definition *discordgo.ApplicationCommand) Command {
	return &command{
		handler:    handler,
		definition: definition,
	}
}
