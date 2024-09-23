package reactions

import (
	"github.com/bwmarrin/discordgo"
	cache "github.com/dabi-ngin/discgo-bot/Cache"
	config "github.com/dabi-ngin/discgo-bot/Config"
	discord "github.com/dabi-ngin/discgo-bot/Discord"
	logger "github.com/dabi-ngin/discgo-bot/Logger"
)

func EditReactions(i *discordgo.InteractionCreate, correlationId string) {
	// 1. Create a Private Thread to handle the options
	channel, err := discord.CreateAdminChannel(i.GuildID, "edit-hall-reactions")
	if err != nil {
		discord.SendEmbedFromInteraction(i, "Error", "There was an error processing your request, please try again.")
		cache.InteractionComplete(correlationId)
		return
	}

	// => Add it to the Interaction Cache for future responses
	cache.ActiveInteractions[correlationId].Values.String["tempChannelId"] = channel.ID

	// 2. By this point we can Acknowledge the original Interaction
	discord.SendEmbedFromInteraction(i, "Editing Hall Reactions", "We have opened a Private Channel for you to apply changes.")

	// 3. Post Messages in the Thread
	// => Introduction Message
	msgText := "<@" + i.Member.User.ID + ">\n"
	msgText += "* Add/Remove the reactions below to define what your Upvote/Downvote emojis should be.\n"
	msgText += "* Click [Reset to Default] to reset them to a standard Thumbs Up/Down.\n"
	msgText += "* Click [Assign ALL Generics] to assign ALL positive/negative default Emojis, these can be edited after.\n"
	msgText += "* Click [Save Changes] to do just that!\n"
	msgText += "* Click [Cancel Changes] discard everything here and destroy this channel.\n"
	_, err = config.Session.ChannelMessageSend(channel.ID, msgText)
	if err != nil {
		discord.SendEmbedFromInteraction(i, "Error", "There was an error processing your request, please try again.")
		cache.InteractionComplete(correlationId)
		discord.DeleteAdminChannel(i.GuildID, channel.ID)
		logger.Error(i.GuildID, err)
		return
	}
	// => Upvote Message
	upMsg, err := config.Session.ChannelMessageSend(channel.ID, "Upvote Reactions")
	if err != nil {
		discord.SendEmbedFromInteraction(i, "Error", "There was an error processing your request, please try again.")
		cache.InteractionComplete(correlationId)
		discord.DeleteAdminChannel(i.GuildID, channel.ID)
		logger.Error(i.GuildID, err)
		return
	}
	// => Downvote Message
	downMsg, err := config.Session.ChannelMessageSend(channel.ID, "Downvote Reactions")
	if err != nil {
		discord.SendEmbedFromInteraction(i, "Error", "There was an error processing your request, please try again.")
		cache.InteractionComplete(correlationId)
		discord.DeleteAdminChannel(i.GuildID, channel.ID)
		logger.Error(i.GuildID, err)
		return
	}

	// 4. Add the Current reactions to these Messages
	for _, emoji := range cache.ActiveGuilds[i.GuildID].ReactionEmojis {
		var emojiErr error
		switch emoji.CategoryID {
		case EmojiCategoryUp:
			emojiErr = config.Session.MessageReactionAdd(channel.ID, upMsg.ID, emoji.Emoji)
		case EmojiCategoryDown:
			emojiErr = config.Session.MessageReactionAdd(channel.ID, downMsg.ID, emoji.Emoji)
		default:
			logger.ErrorText(i.GuildID, "Undefined Emoji category [%v]", emoji.CategoryID)
		}
		if emojiErr != nil {
			logger.Error(i.GuildID, err)
		}
	}

	// 5. Add Buttons
	btnReset := discord.CreateButton(discordgo.Button{
		Label:    "Reset to Default",
		CustomID: "edit-hall-reset",
		Style:    discordgo.SecondaryButton,
	}, correlationId, config.IO_BOUND_TASK, EditHandler_Reset)
	btnApplyAll := discord.CreateButton(discordgo.Button{
		Label:    "Assign ALL Generics",
		CustomID: "edit-hall-apply-all",
		Style:    discordgo.SecondaryButton,
	}, correlationId, config.IO_BOUND_TASK, EditHandler_Generics)
	btnSave := discord.CreateButton(discordgo.Button{
		Label:    "Save Changes",
		CustomID: "edit-hall-save",
		Style:    discordgo.PrimaryButton,
	}, correlationId, config.IO_BOUND_TASK, EditHandler_Save)
	btnCancel := discord.CreateButton(discordgo.Button{
		Label:    "Cancel Changes",
		CustomID: "edit-hall-cancel",
		Style:    discordgo.DangerButton,
	}, correlationId, config.IO_BOUND_TASK, EditHandler_Cancel)

	buttonMessage := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					btnReset,
					btnApplyAll,
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					btnSave,
					btnCancel,
				},
			},
		},
	}

	_, err = config.Session.ChannelMessageSendComplex(channel.ID, buttonMessage)
	if err != nil {
		discord.SendEmbedFromInteraction(i, "Error", "There was an error processing your request, please try again.")
		cache.InteractionComplete(correlationId)
		discord.DeleteAdminChannel(i.GuildID, channel.ID)
		logger.Error(i.GuildID, err)
		return
	}
}

func EditHandler_Reset(i *discordgo.InteractionCreate, correlationId string) {
	logger.Info(i.GuildID, "EditHandler_Reset")
}

func EditHandler_Generics(i *discordgo.InteractionCreate, correlationId string) {
	logger.Info(i.GuildID, "EditHandler_Generics")
}

func EditHandler_Save(i *discordgo.InteractionCreate, correlationId string) {
	logger.Info(i.GuildID, "EditHandler_Save")
}

func EditHandler_Cancel(i *discordgo.InteractionCreate, correlationId string) {
	logger.Info(i.GuildID, "EditHandler_Cancel")
}
