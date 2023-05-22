package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/zVasO/StalkerCypher/config"
	"log"
)

func getCommands() []*discordgo.ApplicationCommand {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "play",
			Description: "Play music",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "query",
					Description: "the youtube url",
					Required:    true,
				},
			},
		},
	}
	return commands
}

func registerCommands(goBot *discordgo.Session) {
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
