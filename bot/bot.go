package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/ErneDainesi/discord-music-bot/yt"
	"github.com/bwmarrin/discordgo"
)

func Run(token string) {
    discord, err := discordgo.New("Bot " + token)
    if err != nil {
        fmt.Print("hubo un error\n")
        return
    }

    discord.AddHandler(manageMessage)

    discord.Open()
    defer discord.Close()
    fmt.Println("Bot running....")
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
}

func manageMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	/* prevent bot responding to its own message
	   this is achived by looking into the message author id
	   if message.author.id is same as bot.author.id then just return
	*/
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// respond to user message if it contains `!help` or `!bye`
	switch {
    case strings.Contains(message.Content, "!play"):
        err := yt.Download(message.Content)
        if err != nil {
            discord.ChannelMessageSend(message.ChannelID, "Failed to download video")
            fmt.Println("Failed to download video")
            return
        }
	}
}
