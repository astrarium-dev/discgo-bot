package bang

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	config "github.com/dabi-ngin/discgo-bot/Config"
	database "github.com/dabi-ngin/discgo-bot/Database"
	helpers "github.com/dabi-ngin/discgo-bot/Helpers"
)

type DelImage struct{}

func (s DelImage) Name() string {
	return "DelImage"
}

func (s DelImage) PermissionRequirement() int {
	return config.CommandLevelUser
}

func (s DelImage) ProcessPool() config.ProcessPool {
	return config.ProcessPools[config.ProcessPoolText]
}

func (s DelImage) LockedByDefault() bool {
	return true
}

func (s DelImage) Execute(message *discordgo.MessageCreate, command string) error {

	imgUrl := helpers.GetImageFromMessage(message.Message, "")
	if imgUrl == "" {
		return errors.New("no image found")
	}

	imgCat, err := database.GetImgCategory(message.GuildID, command)
	if err != nil {
		return errors.New("unable to get gif category")
	}

	imgStorage, err := database.GetImgStorage(message.GuildID, imgUrl)
	if err != nil {
		return errors.New("unable to get gif category")
	}

	imgGuildLink, err := database.GetImgGuildLink(message.GuildID, imgCat, imgStorage)
	if err != nil {
		return errors.New("unable to get guild link")
	}

	return database.DeleteGuildLink(imgGuildLink)

}
