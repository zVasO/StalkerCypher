package commands

import (
	"StalkerCypher/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

func getCommands() []*discordgo.ApplicationCommand {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "rank",
			Description: "Give the rank",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "query",
					Description: "Username #TAG",
					Required:    true,
				},
			},
		},
	}
	return commands
}

func RegisterCommands(goBot *discordgo.Session) {
	commands := getCommands()

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := goBot.ApplicationCommandCreate(config.AppID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}
