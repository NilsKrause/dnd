package character

import (
	"de.nilskrau.dndbot/command"
	"github.com/bwmarrin/discordgo"
)

var DeleteCharacterCommand = command.NewCommand(
	func(cc command.Context) {
		chars, err := cc.GetUserCharacters()
		if err != nil {
			cc.SendErrorMessage("Konnte deine Charaktere nicht laden :(")
			return
		}

		id := cc.GetOption("id")
		if id == nil {
			cc.SendErrorMessage("Eine id ist nötig zum löschen")
			return
		}

		for _, char := range chars {
			if char.ID == uint(id.IntValue()) {
				cc.GetApp().DeleteCharacter(char)
				cc.SendInfo("Charakter gelöscht")
				return
			}
		}
	
		cc.SendErrorMessage("Konnte Charakter mit ID nicht finden :(")
	},
	&discordgo.ApplicationCommand{
		Name:        "delete",
		Description: "Character löschen",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name: "id",
				Description: "Die id des Charakters, zu entnehmen aus dem /list kommando",
				Type: discordgo.ApplicationCommandOptionInteger,
				Required: true,
			},
		},
	},
)
