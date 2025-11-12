package write

import (
	"fmt"

	"de.nilskrau.dndbot/command"
	"de.nilskrau.dndbot/models"
	"github.com/bwmarrin/discordgo"
)

var WriteIncharacterCommand = command.NewCommand(
	func(cc command.Context) {
		var char *models.Character
		var err error
		handle := cc.GetOption("handle")
		if handle == nil {
			char, err = cc.GetUserActiveCharacter()
			if err != nil {
				cc.SendErrorMessage("Kein Charakter ist als Aktiv markiert.")
				fmt.Println(err)
				return
			}
		} else {
			if handle.Type != discordgo.ApplicationCommandOptionString {
				cc.SendErrorMessage("Handle muss ein text sein.")
				fmt.Println(err)
				return
			}
			char, err = cc.GetUserCharacterByHandle(handle.StringValue())
			if err != nil {
				cc.SendErrorMessage("Konnte Charakter mit handle " + handle.StringValue() + " nicht finden.")
				fmt.Println(err)
				return
			}
		}

		if char == nil {
			chars, err := cc.GetUserCharacters()
			if err != nil {
				cc.SendErrorMessage("Konnte deine Charaktere nicht von der Datenbank laden :(")
				fmt.Println(err)
			}
			if len(chars) == 1 {
				char = chars[0]
			} else {
				cc.SendErrorMessage("Du hast leider noch keine Charaktere angelegt :( - lege einen Char mit /character an!")
				return
			}
		}

		player, err := cc.GetPlayer()
		if err != nil {
			cc.SendErrorMessage("Konnte dich nicht als Spieler identifizieren o.O")
			fmt.Println(err)
			return
		}

		msg := cc.GetOption("message")
		if msg == nil {
			cc.SendErrorMessage("Du solltest auch eine Nachricht schreiben... sonst bringt das ganze nix :D")
			return
		} else if msg.Type != discordgo.ApplicationCommandOptionString {
			cc.SendErrorMessage("Deine Nachricht muss text sein")
			return
		}
		
		cc.GetApp().LogMessage(player, char, msg.StringValue())
		err = cc.SendIncharacterEmbed(char, player, msg.StringValue())
		if err != nil {
			cc.SendErrorMessage("Konnte die Nachricht nicht senden :(")
			fmt.Println(err)
			return
		}

		err = cc.SendInfo("Nachricht versandt")
	},
	&discordgo.ApplicationCommand{
		Name:        "write",
		Description: "Eine Nachricht in Charakter verfassen",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "message",
				Description: "Die Nachricht die du verfassen möchtest.",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},

			{
				Name:        "handle",
				Description: "Handle des Charakters in dem du schreiben möchtest.",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    false,
			},
		},
	},
)
