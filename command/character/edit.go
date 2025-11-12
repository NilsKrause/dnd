package character

import (
	"fmt"
	"net/url"

	"de.nilskrau.dndbot/command"
	"github.com/bwmarrin/discordgo"
)

var EditCharacterCommand = command.NewCommand(
	func(cc command.Context) {
		userid := cc.GetUser().ID

		p, err := cc.GetApp().CreateOrGetPlayer(userid)
		if err != nil {
			cc.SendErrorMessage("Konnte dich nicht in meiner Datenbank finden :(")
			fmt.Println(err)
			return
		}

		id := cc.GetOption("id")
		if id == nil || id.Type != discordgo.ApplicationCommandOptionInteger {
			cc.SendErrorMessage("Die id muss eine Zahl sein.")
			fmt.Println("id is nil or not an int")
			return
		}

		name := cc.GetOption("name")
		if name == nil || name.Type != discordgo.ApplicationCommandOptionString {
			cc.SendErrorMessage("Dein Name muss Text sein.")
			fmt.Println("Name nil or not a string")
			return
		}

		handle := cc.GetOption("handle")
		if handle == nil || handle.Type != discordgo.ApplicationCommandOptionString {
			cc.SendErrorMessage("Das handle muss Text sein.")
			fmt.Println("handle nil or not a string")
			return
		}

		picture := cc.GetOption("picture")
		if picture == nil || picture.Type != discordgo.ApplicationCommandOptionString {
			cc.SendErrorMessage("Es muss ein Bild angegeben werden.")
			fmt.Println("picture is nil or not a string")
			return
		}
		_, err = url.Parse(picture.StringValue())
		if err != nil {
			cc.SendErrorMessage("Du musst einen Öffentlichen URL zu seinem Bild angeben")
			fmt.Println("picture was not a url", err)
			return
		}

		activeOpt := cc.GetOption("active")
		if activeOpt != nil && activeOpt.Type != discordgo.ApplicationCommandOptionBoolean {
			cc.SendErrorMessage("Active muss entweder 'True' oder 'False' sein")
			fmt.Println("active not a bool")
			return
		}
		active := false
		if activeOpt != nil {
			active = activeOpt.BoolValue()
		}

		_, err = cc.GetApp().EditCharacter(uint(id.IntValue()), p, name.StringValue(), handle.StringValue(), picture.StringValue(), active)
		if err != nil {
			cc.SendErrorMessage("Konnte deinen character leider nicht editieren :(")
			fmt.Println("could not edit character", err)
			return
		}

		if err = cc.SendInfo("Character erfolgreich bearbeitet"); err != nil {
			fmt.Println("error while sending success message", err)
		}
	},
	&discordgo.ApplicationCommand{
		Name:        "edit",
		Description: "Character bearbeiten",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name: 			 "id",
				Description: "Die ID des Charakters, zu entnehmen aus /list",
				Type: 			 discordgo.ApplicationCommandOptionInteger,
				Required: 	 true,
			},
			{
				Name:        "name",
				Description: "Name der im Chat angezeigt wird",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "handle",
				Description: "Abkürzung zur Verwendung im Chat",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "picture",
				Description: "Öffentlicher URL zum Bild, kann URL zu einem Bild aus einem Discord Chat sein",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "active",
				Description: "Ob dieser Character als Standard verwendet werden soll",
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Required:    false,
			},
		},
	},
)
