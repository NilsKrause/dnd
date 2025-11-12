package character

import (
	"fmt"
	"net/url"

	"de.nilskrau.dndbot/command"
	"github.com/bwmarrin/discordgo"
)

var CreateCharacterCommand = command.NewCommand(
	func(cc command.Context) {
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

		chars, err := cc.GetUserCharacters()
		if err != nil {
			cc.SendErrorMessage("Konnte nicht prüfen ob die handle einzigartig ist")
			fmt.Println("error while loading chars", err)
			return
		}

		for _, c := range chars {
			if c.Handle == handle.StringValue() {
				cc.SendErrorMessage("Handle ist nicht einzigartig.")
				fmt.Println("handle is not unique")
				return
			}
		}

		picture := cc.GetOption("picture")
		if picture == nil || picture.Type != discordgo.ApplicationCommandOptionString {
			cc.SendErrorMessage("Es muss ein Bild angegeben werden.")
			fmt.Println("picture is nil or not a string")
			return
		}
		_, err = url.ParseRequestURI(picture.StringValue())
		if err != nil {
			cc.SendErrorMessage("Du musst einen öffentlichen URL zu einem Bild angeben")
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

		err = cc.CreateCharacter(name.StringValue(), handle.StringValue(), picture.StringValue(), active)
		if err != nil {
			cc.SendErrorMessage("Konnte deinen character leider nicht erstellen :(")
			fmt.Println("could not create character", err)
			return
		}

		if err = cc.SendInfo("Character erfolgreich erstellt"); err != nil {
			fmt.Println("error while sending success message", err)
		}
	},
	&discordgo.ApplicationCommand{
		Name:        "character",
		Description: "Character erstellen",
		Options: []*discordgo.ApplicationCommandOption{
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
