package main

import (
	"fmt"
	"github.com/zVasO/StalkerCypher/bot"
	"github.com/zVasO/StalkerCypher/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
