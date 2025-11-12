package command

import (
	"errors"
	"fmt"
	"log"

	"de.nilskrau.dndbot/models"
	"github.com/bwmarrin/discordgo"
)

type commandContext struct {
	BotContext
	options     Options
	interaction *discordgo.InteractionCreate
}

func (cc *commandContext) GetPlayer() (*models.Player, error) {
	return cc.GetApp().CreateOrGetPlayer(cc.GetUser().ID)
}

func (cc *commandContext) GetUserCharacterByHandle(handle string) (*models.Character, error) {
	chars, err := cc.GetUserCharacters()
	if err != nil {
		return nil, err
	}

	for _, c := range chars {
		if c.Handle == handle {
			return c, nil
		}
	}

	return nil, errors.New("character with handle not found")
}

func (cc *commandContext) GetUserActiveCharacter() (*models.Character, error) {
	chars, err := cc.GetUserCharacters()
	if err != nil {
		return nil, err
	}

	for _, c := range chars {
		if c.Default {
			return c, nil
		}
	}

	return nil, errors.New("no active character found")
}

func (cc *commandContext) GetUserCharacters() ([]*models.Character, error) {
	p, err := cc.GetPlayer()
	if err != nil {
		return nil, err
	}

	chars, err := cc.GetApp().GetPlayerCharacters(p, cc.GetCurrentServerId())
	if err != nil {
		return nil, err
	}

	return chars, nil
}

func (cc *commandContext) GetUser() *discordgo.User {
	i := cc.interaction
	if i.User != nil {
		return i.User
	} else if i.Member != nil {
		return i.Member.User
	}

	code := "neigher user or member.user are set"

	log.Print(code)
	panic(code)
}

func (cc *commandContext) SendErrorMessage(message string) error {
	return cc.GetSession().InteractionRespond(cc.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Content: message,
		},
	})
}

func (cc *commandContext) SendCharacterEmbeds(chars []*models.Character) error {
	embeds := make([]*discordgo.MessageEmbed, 0)
	for _, char := range chars {
		text := fmt.Sprintf("ID: %d; Handle: %s", char.ID, char.Handle)
		if char.Default {
			text = fmt.Sprintf("ID: %d; Handle: %s; Active", char.ID, char.Handle)
		}
		embeds = append(embeds, &discordgo.MessageEmbed{
			Title:     char.Name,
			//Image: &discordgo.MessageEmbedImage{URL: char.Image},
			Thumbnail: &discordgo.MessageEmbedThumbnail{URL: char.Image},
			Footer: &discordgo.MessageEmbedFooter{
				Text: text,
			},
		})
	}

	return cc.GetSession().InteractionRespond(cc.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: embeds,
		},
	})

}

func (cc *commandContext) SendIncharacterEmbed(char *models.Character, player *models.Player, message string) error {
	user := cc.GetUser()
	embed := &discordgo.MessageEmbed{
		Title:     char.Name + ":",
		Description: message,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: char.Image},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "@" + user.GlobalName,
			IconURL: user.AvatarURL(""),
		},
	}

	_, err := cc.GetSession().ChannelMessageSendEmbed(cc.interaction.ChannelID, embed)
	if err != nil {
		fmt.Println(err)
		return err
	}
	

	return nil
}

func (cc *commandContext) SendInfo(message string) error {
	return cc.GetSession().InteractionRespond(cc.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Content: message,
		},
	})
}

func (cc *commandContext) SendMessage(text string) error {
	return cc.GetSession().InteractionRespond(cc.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: text,
		},
	})
}

func (cc *commandContext) GetOptions() Options {
	return cc.options
}

func (cc *commandContext) GetOption(option string) *discordgo.ApplicationCommandInteractionDataOption {
	return cc.options[option]
}

func (cc *commandContext) CreateCharacter(name string, handle string, picture string, active bool) error {
	p, err := cc.GetPlayer()
	if err != nil {
		return err
	}

	_, err = cc.GetApp().CreateCharacter(p, cc.GetCurrentServerId(), name, handle, picture, active)
	if err != nil {
		return err
	}

	return nil
}

func (cc *commandContext) GetCurrentServerId() string {
	return cc.interaction.GuildID
}

func buildCommandContext(botCtx BotContext, i *discordgo.InteractionCreate) Context {
	data := i.ApplicationCommandData()
	o := make(Options)
	for _, opt := range data.Options {
		o[opt.Name] = opt
	}

	return &commandContext{
		BotContext:  botCtx,
		options:     o,
		interaction: i,
	}
}

