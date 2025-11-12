package character

import (
	"de.nilskrau.dndbot/command"
	"github.com/bwmarrin/discordgo"
)

var ListCharacterCommand = command.NewCommand(
	func(cc command.Context) {
		chars, err := cc.GetUserCharacters()
		if err != nil {
			cc.SendErrorMessage("Konnte deine Charaktere nicht laden :(")
			return
		}

		if len(chars) == 0 {
			cc.SendInfo("Du hast leider noch keinen Charakter, erstelle einen mit /character")
		}
	
		err = cc.SendCharacterEmbeds(chars)
		if err != nil {
			cc.SendErrorMessage("Konnte dir die Liste nicht zusenden :(")
		}
		
	},
	&discordgo.ApplicationCommand{
		Name:        "list",
		Description: "Charactere auflisten",
		Options: []*discordgo.ApplicationCommandOption{},
	},
)
