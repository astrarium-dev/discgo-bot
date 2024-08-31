package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// These can be swapped around on the go, but pls don't lol. If any are added make sure to also update the map
const (
	LoggingLevelAdmin = iota
	LoggingLevelError
	LoggingLevelWarn
	LoggingLevelEvent
	LoggingLevelInfo
	LoggingLevelDebug
)

var LoggingLevels map[int]LoggingOptions = map[int]LoggingOptions{
	LoggingLevelAdmin: {
		Name:   "Admin",
		Colour: Colours["magenta"],
	},
	LoggingLevelError: {
		Name:   "Error",
		Colour: Colours["red"],
	},
	LoggingLevelWarn: {
		Name:   "Warn",
		Colour: Colours["yellow"],
	},
	LoggingLevelEvent: {
		Name:   "Event",
		Colour: Colours["green"],
	},
	LoggingLevelInfo: {
		Name:   "Info",
		Colour: Colours["white"],
	},
	LoggingLevelDebug: {
		Name:   "Debug",
		Colour: Colours["default"],
	},
}

type LoggingOptions struct {
	Name   string
	Colour Colour
}

type Colour struct {
	Terminal string `json:"Terminal,omitempty"`
	Html     string `json:"Html,omitempty"`
}

var Colours map[string]Colour = map[string]Colour{
	"default": {
		Terminal: "\033[0m",
		Html:     "#000000",
	},
	"white": {
		Terminal: "\033[97m",
		Html:     "#FFFFFF",
	},
	"magenta": {
		Terminal: "\033[35m",
		Html:     "#C30CC9",
	},
	"yellow": {
		Terminal: "\033[33m",
		Html:     "#FAF200",
	},
	"green": {
		Terminal: "\033[32m",
		Html:     "#28F200",
	},
	"red": {
		Terminal: "\033[31m",
		Html:     "#F20008",
	},
}

const (
	CommandLevelBotAdmin = iota
	CommandLevelServerOwner
	CommandLevelAdmin
	CommandLevelMod
	CommandLevelVIP
	CommandLevelUser
	CommandLevelRestricted
	CommandLevelDisabled
)

// Task categories for channels in message handling
const (
	// TRIVIAL_TASK involves small CPU and no IO waiting
	TRIVIAL_TASK = iota
	// CPU_BOUND_TASK involves intensive operations
	CPU_BOUND_TASK
	// IO_BOUND_TASK involves waiting on API/DB response
	IO_BOUND_TASK
)

const (
	N_TRIVIAL_WORKERS = 50
	N_IO_WORKERS      = 5
)

// Command Types
// This is used to denote types to the Dashboard
// ------------------------------------------------
const ( // A) Populate A + B!
	CommandTypeDefault = iota
	CommandTypeBang
	CommandTypePhrase
)

// B) Populate A + B!
var CommandTypes map[int]string = map[int]string{
	CommandTypeDefault: "Default",
	CommandTypeBang:    "Bang",
	CommandTypePhrase:  "Phrase",
}

// ------------------------------------------------

// Process Pools
// Used to dispatch BangCommands in the newMessage Handler
// ------------------------------------------------
const (
	ProcessPoolText = iota
	ProcessPoolImages
	ProcessPoolExternal
)

var LastPoolIota int = ProcessPoolExternal

var ProcessPools map[int]ProcessPool = map[int]ProcessPool{
	ProcessPoolText: {
		ProcessPoolIota: ProcessPoolText,
		PoolName:        "Text",
		MaxWorkers:      50,
	},
	ProcessPoolImages: {
		ProcessPoolIota: ProcessPoolImages,
		PoolName:        "Images",
		MaxWorkers:      25,
	},
	ProcessPoolExternal: {
		ProcessPoolIota: ProcessPoolExternal,
		PoolName:        "External",
		MaxWorkers:      10,
	},
}

// -------------------------------------------------
type ProcessPool struct {
	ProcessPoolIota int
	PoolName        string
	MaxWorkers      int
}

type Vars struct {
	IsDev       bool
	SuperAdmins []string

	LogToDiscord       bool
	LoggingChannelID   string
	LoggingUsesThreads bool
	VerboseStack       bool
	LogFunctions       bool
	LoggingLevel       int

	DashboardMaxDataPackets     int
	DashboardMaxLogs            int
	DashboardMaxCommands        int
	CommandAveragePool          int
	HardwareStatIntervalSeconds int
	HardwareStatMaxIntervals    int

	BotToken string

	DB_NAME       string
	DB_USER       string
	DB_PASSWORD   string
	DB_IP_ADDRESS string
	DB_PORT       string
}

var (
	IsDev        bool
	HostName     string
	SuperAdmins  []string
	DashboardUrl string

	LogToDiscord        bool
	LoggingChannelID    string
	LoggingUsesThreads  bool
	LoggingVerboseStack bool
	LoggingLogFunctions bool

	DashboardMaxDataPackets     int
	DashboardMaxLogs            int
	DashboardMaxCommands        int
	CommandAveragePool          int
	HardwareStatIntervalSeconds int
	HardwareStatMaxIntervals    int

	LoggingLevel int

	BotToken string
	Session  *discordgo.Session

	ValidImageExtensions []string

	DB_NAME       string
	DB_USER       string
	DB_PASSWORD   string
	DB_IP_ADDRESS string
	DB_PORT       string
)

const (
	MAX_MESSAGE_LENGTH int    = 2000
	ROOT_FOLDER        string = "discgo-bot/"
	BOT_SUB_FOLDER     string = "Bot/"
)

func Init() bool {
	localConfigFile, err := os.ReadFile("config.json")

	if err != nil {
		fmt.Println(fmt.Printf("Config.Init() - Error loading config.json :: %v", err))
		return false
	}

	var configFileVariables Vars
	err = json.Unmarshal([]byte(localConfigFile), &configFileVariables)
	if err != nil {
		fmt.Println(fmt.Printf("Config.Init() - Error unmarshalling config.json :: %v", err))
		return false
	}

	currentHostName, err := os.Hostname()
	if err != nil {
		HostName = "Unknown"
	} else {
		HostName = currentHostName
	}

	ValidImageExtensions = []string{
		".gif",
		".png",
		".jpg",
		".webp",
	}

	IsDev = configFileVariables.IsDev
	if IsDev {
		DashboardUrl = "http://localhost:3333/"
	}

	SuperAdmins = configFileVariables.SuperAdmins

	LoggingChannelID = configFileVariables.LoggingChannelID
	LoggingUsesThreads = configFileVariables.LoggingUsesThreads
	LoggingVerboseStack = configFileVariables.VerboseStack
	LoggingLogFunctions = configFileVariables.LogFunctions
	LoggingLevel = configFileVariables.LoggingLevel

	DashboardMaxDataPackets = configFileVariables.DashboardMaxDataPackets
	DashboardMaxLogs = configFileVariables.DashboardMaxLogs
	DashboardMaxCommands = configFileVariables.DashboardMaxCommands
	CommandAveragePool = configFileVariables.CommandAveragePool
	HardwareStatIntervalSeconds = configFileVariables.HardwareStatIntervalSeconds
	HardwareStatMaxIntervals = configFileVariables.HardwareStatMaxIntervals
	LogToDiscord = configFileVariables.LogToDiscord

	BotToken = configFileVariables.BotToken

	DB_NAME = configFileVariables.DB_NAME
	DB_USER = configFileVariables.DB_USER
	DB_PASSWORD = configFileVariables.DB_PASSWORD
	DB_IP_ADDRESS = configFileVariables.DB_IP_ADDRESS
	DB_PORT = configFileVariables.DB_PORT

	return true
}
