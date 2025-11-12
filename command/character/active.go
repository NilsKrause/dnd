package character

import (
	"fmt"

	"de.nilskrau.dndbot/command"
	"github.com/bwmarrin/discordgo"
)

var SetActiveCharacterCommand = command.NewCommand(
	func(cc command.Context) {
		p, err := cc.GetPlayer()
		if err != nil {
			cc.SendErrorMessage("Konnte dich nicht in der Datenbank finden :(")
			fmt.Println("error looking up player in db")
			return
		}

		chars, err := cc.GetUserCharacters()
		if err != nil {
			cc.SendErrorMessage("Konnte deine Charaktere nicht laden :(")
			return
		}

		if len(chars) == 0 {
			cc.SendInfo("Du hast leider noch keinen Charakter, erstelle einen mit /character")
			return
		}

		id := cc.GetOption("id")
		if id == nil || id.Type != discordgo.ApplicationCommandOptionInteger {
			cc.SendErrorMessage("Id muss eine Nummer sein.")
			fmt.Println("id is nil or not an int")
			return
		}
	
		char, err := cc.GetApp().CharacterSetActive(uint(id.IntValue()), p, true)
		if err != nil {
			cc.SendErrorMessage("Konnte Charakter nicht aktiv setzen")
			fmt.Println("error while setting char active", err)
			return
		}

		for _, char := range chars {
			if !char.Default {
				continue
			}
			_, err = cc.GetApp().CharacterSetActive(char.ID, p, false)
			if err != nil {
				cc.SendErrorMessage("Konnte aktiven Charakter nicht deaktivieren")
				fmt.Println("error while deactivating character", err)
				return
			}
		}

		cc.SendInfo("Charakter " + char.Name + " erfolgreich aktiv gesetzt")
	},
	&discordgo.ApplicationCommand{
		Name:        "activate",
		Description: "Charactere aktiv setzen",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name: "id",
				Description: "Die id des Charakters den du aktivieren möchtest",
				Type: discordgo.ApplicationCommandOptionInteger,
				Required: true,
			},
		},
	},
)
