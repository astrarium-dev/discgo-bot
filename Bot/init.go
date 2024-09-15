package bot

import (
	"errors"
	"os"

	"github.com/bwmarrin/discordgo"
	handlers "github.com/dabi-ngin/discgo-bot/Bot/Handlers"
	config "github.com/dabi-ngin/discgo-bot/Config"
	dashboard "github.com/dabi-ngin/discgo-bot/Dashboard"
	database "github.com/dabi-ngin/discgo-bot/Database"
	logger "github.com/dabi-ngin/discgo-bot/Logger"
	"github.com/google/uuid"
)

func Init() {
	// 1. Database Init
	if !database.Init() {
		logger.Error("", errors.New("Failed to initialise database"))
		return
	}

	// 2. Discord Session Init
	if !sessionInit() {
		logger.Error("", errors.New("Failed to initialise session"))
		return
	}

	// 3. Add Handlers to the Session
	if !addHandlers() {
		logger.Error("", errors.New("Failed to add handlers"))
		return
	}

	// 4. Open the Discord session
	if !sessionOpen() {
		logger.Error("", errors.New("Failed to open session"))
		return
	}

	// 5. Log Init
	if !logger.Init() {
		logger.Error("", errors.New("Failed to initialise logging"))
		return
	}

	var initSuffix string
	if config.ServiceSettings.DASHBOARDURL != "" {
		initSuffix = ", Dashboard open at: " + config.ServiceSettings.DASHBOARDURL
	}
	logger.Info("", "Bot intialisation successful%s", initSuffix)

	if !config.ServiceSettings.ISDEV {
		dashboard.TrySendDashboardInitMessage()
	}

	// 6. Register Discord /commands
	if !registerCommands() {
		logger.Error("", errors.New("Failed to register commands"))
		return
	}

	// 7. Reset Global Discord /commands
	handlers.RefreshSlashCommands("")

	// -------------------
	file, err := os.Open("87673673-6626-49a9-b9b2-944b1ba90e1d.wav") // replace with your file name
	if err != nil {
		logger.Error("TEST", err)
		return
	}
	defer file.Close()

	// Create discordgo.File object
	fileObj := &discordgo.File{
		Name:   uuid.New().String() + ".wav",
		Reader: file,
	}
	// Send the message with embed and file
	outputText := "-# 'Senator' by <@192015008039698432>, reply with **!ttsinfo** for details"
	_, err = config.Session.ChannelMessageSendComplex("1269346890147827774", &discordgo.MessageSend{
		Content: outputText,
		Files:   []*discordgo.File{fileObj},
	})
	if err != nil {
		logger.Error("TEST", err)
	}
	// -------------------
}

func sessionInit() bool {
	session, err := discordgo.New("Bot " + config.ServiceSettings.BOTTOKEN)
	if err != nil {
		logger.Error("", err)
		return false
	} else if session == nil {
		logger.Error("", err)
		return false
	}

	config.Session = session
	return true
}

func sessionOpen() bool {
	err := config.Session.Open()
	if err != nil {
		logger.Error("", err)
		return false
	}

	if config.Session == nil {
		logger.Error("FUCK", err)
	}
	return true
}

func addHandlers() bool {
	config.Session.AddHandler(handlers.HandleNewMessage)          //   New Messages
	config.Session.AddHandler(handlers.HandleNewGuild)            //	Server connected to the bot
	config.Session.AddHandler(handlers.HandleInteractionResponse) //	Responses from Interaction objects
	return true
}

func registerCommands() bool {
	return true
}
