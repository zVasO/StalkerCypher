package main

import (
	"fmt"
	"github.com/DylanGermann/DjNoopy/bot"
	"github.com/DylanGermann/DjNoopy/config"
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
