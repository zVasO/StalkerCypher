package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/zVasO/StalkerCypher/config"
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
	registerCommands(goBot)

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
		options := i.ApplicationCommandData().Options

		// Or convert the slice into a map
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		query := optionMap["query"].StringValue()

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Hey <@%s> As u asked, im gonna play %s", i.Interaction.Member.User.ID, query),
			},
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func Play() {

}
