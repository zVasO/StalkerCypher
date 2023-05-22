package commands

import (
	"StalkerCypher/riot"
	"fmt"
	riot2 "github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/clients/val"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Rank(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	query := optionMap["query"].StringValue()
	user := getUserAccount(query)
	fmt.Println(user)
	temp := getRank(user.PUUID)
	fmt.Println(temp)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Hey <@%s> As u asked, im gonna show you the rank of %s", i.Interaction.Member.User.ID, query),
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getRank(userId string) string {
	client, err := riot.GetRiotClient()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client, userId)

	list, err := client.VAL.Match.List(val.EU, userId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
	//exportJson(list.History)
	//fmt.Println(list.History[0])
	return ""
}

func getCurrentAct(contents *val.LocalizedContentDTO) val.LocalizedActDTO {
	for _, act := range contents.Acts {
		if act.IsActive {
			return act
		}
	}
	panic("No act found")
}

func getAllActs(contents *val.LocalizedContentDTO) []val.LocalizedActDTO {
	return contents.Acts
}

func getAllContentsByLocal(shard val.Shard, locale val.Locale) *val.LocalizedContentDTO {
	client, err := riot.GetRiotClient()
	if err != nil {
		fmt.Println(err)
	}

	allContent, err := client.VAL.Content.ByLocale(shard, locale)
	if err != nil {
		fmt.Println(err.Error())
	}
	return allContent
}

func getUserAccount(username string) *riot2.AccountDTO {
	client, err := riot.GetRiotClient()
	if err != nil {
		fmt.Println(err)
	}
	splitUsername := verifyUsername(username)

	userAccount, err := client.Riot.Account.ByID(splitUsername[0], splitUsername[1])
	if err != nil {
		fmt.Println(err)
	}
	return userAccount
}
func verifyUsername(username string) []string {
	isValid := strings.Contains(username, "#")
	if !isValid {
		fmt.Printf("The username %s does not contains a tag (ex: #EUW)", username)
	}

	return strings.Split(username, "#")
}

func exportJson(data string) {
	file, err := os.Create("data.json")
	if err != nil {
		return
	}
	write, err := file.Write([]byte(data))
	if err != nil {
		return
	}
	fmt.Println(write)
}
