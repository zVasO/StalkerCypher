package bot

import (
	"StalkerCypher/commands"
	"StalkerCypher/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = u.ID
	commands.RegisterCommands(goBot)

	goBot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	go Play()
	fmt.Println("Bot is running !")
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"rank": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		commands.Rank(s, i)
	},
}

func Play() {

}
