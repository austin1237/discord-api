package main

import (
	"fmt"
	"os"

	"github.com/user/discordapi/discord"
)

var (
	//ENV VARIABLES
	discordServerID string
	discordSecret   string
)

func init() {
	discordSecret = os.Getenv("DISCORD_SECRET")
	if discordSecret == "" {
		fmt.Println("DISCORD_SECRET ENV var was not set.")
		os.Exit(1)
	}

	discordServerID = os.Getenv("DISCORD_ID")
	if discordServerID == "" {
		fmt.Println("DISCORD_ID ENV var was not set.")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Hello, from main!")
	message := "json marshaling"
	err := discord.SendMessage(discordServerID, discordSecret, message)
	if err != nil {
		fmt.Println("error occured" + err.Error())
		os.Exit(1)
	}
}
