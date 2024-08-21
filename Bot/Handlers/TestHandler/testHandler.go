package testhandler

import (
	"github.com/bwmarrin/discordgo"
	config "github.com/dabi-ngin/discgo-bot/Config"
	logger "github.com/dabi-ngin/discgo-bot/Logger"
	structs "github.com/dabi-ngin/discgo-bot/Structs"
)

func HandleNewMessage(message *discordgo.MessageCreate, self structs.BangCommand) error {
	_, err := config.Session.ChannelMessageSendReply(message.ChannelID, "Test", message.Reference())

	if err != nil {
		logger.Error(message.GuildID, err)
	}

	return err
}

func HandleNewTrigger(message *discordgo.MessageCreate, trigger string) error {
	_, err := config.Session.ChannelMessageSendReply(message.ChannelID, "TRIGGER DETECTED", message.Reference())

	if err != nil {
		logger.Error(message.GuildID, err)
	}
	return err
}
