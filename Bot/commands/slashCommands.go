package commands

import (
	"sort"

	"github.com/ZestHusky/femboy-control/Bot/config"
	"github.com/ZestHusky/femboy-control/Bot/constants"
	dbhelper "github.com/ZestHusky/femboy-control/Bot/dbhelpers"
	"github.com/ZestHusky/femboy-control/Bot/friday"
	"github.com/ZestHusky/femboy-control/Bot/gifbank"
	"github.com/ZestHusky/femboy-control/Bot/handlers/animegif"
	"github.com/ZestHusky/femboy-control/Bot/handlers/ask"
	"github.com/ZestHusky/femboy-control/Bot/handlers/dice"
	"github.com/ZestHusky/femboy-control/Bot/handlers/fakeyou"
	"github.com/ZestHusky/femboy-control/Bot/handlers/help"
	"github.com/ZestHusky/femboy-control/Bot/handlers/meme"
	"github.com/ZestHusky/femboy-control/Bot/handlers/slurs"
	"github.com/ZestHusky/femboy-control/Bot/handlers/todo"
	"github.com/ZestHusky/femboy-control/Bot/handlers/userstats"
	"github.com/ZestHusky/femboy-control/Bot/helpers"
	"github.com/ZestHusky/femboy-control/Bot/logging"
	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "fwiday",
		Description: "Is it Fwiday? Check!",
	},
	{
		Name:        "friday",
		Description: "Is it Friday? Check!",
	},
	{
		Name:        "reacts",
		Description: "Get a (private) list of all the available Anime !reaction commands",
	},
	{
		Name:        "bot-help",
		Description: "Get a list of all ! commands available",
	},
	{
		Name:        "voicelist",
		Description: "Get a (private) list of all the available !TTS voices",
	}, {
		Name:        "slur-definition",
		Description: "Make sure you know exactly what connotations are attached to the slurs you haul at people",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "slur",
				Description: "The Slur to Check",
				Required:    true,
			},
		},
	},
	{
		Name:        "todolist",
		Description: "View the ToDo List",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "category",
				Description: "ToDo Category, leave blank to show all",
				Required:    false,
				Choices:     todo.ToDoTypeSelector(),
			},
		},
	},
	{
		Name:        "todoadd",
		Description: "Add a ToDo Item",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "category",
				Description: "ToDo Category",
				Required:    true,
				Choices:     todo.ToDoTypeSelector(),
			}, {
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "item",
				Description: "ToDo Text",
				Required:    true,
			},
		},
	},
	{
		Name:        "todoupdate",
		Description: "Update an existing ToDo item",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "ToDo ID (ie. SWKFEM-1001 or W1001)",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "started",
				Description: "Started by User",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "finished",
				Description: "Finished by User",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "newtext",
				Description: "New ToDo text",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "newcategory",
				Description: "New ToDo Category",
				Required:    false,
				Choices:     todo.ToDoTypeSelector(),
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "version",
				Description: "Released in Version",
				Required:    false,
			},
		},
	},
	{
		Name:        "tododelete",
		Description: "Search FakeYou's TTS models",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "ToDo ID (ie. SWKFEM-1001 or W1001)",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "delete-confirmation",
				Description: "Type 'delete' just to make ABSOLUTELY sure",
				Required:    true,
			},
		},
	},
	{
		Name:        "ttssearch",
		Description: "Search FakeYou's TTS models",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "search",
				Description: "Search Term",
				Required:    true,
			},
		},
	},
	{
		Name:        "ttsadd",
		Description: "Search FakeYou's TTS models",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "command",
				Description: "The text to call the model, !tts[command]",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "description",
				Description: "The description for this model which will appear in /voicelist",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "model",
				Description: "The model id provided from the /ttssearch function",
				Required:    true,
			},
		},
	},
	{
		Name:        "ttsupdate",
		Description: "Update an existing TTS model",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "command",
				Description: "The command to edit (ie. for !ttspetah: petah)",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "new-command",
				Description: "Changes the command if entered",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "new-model",
				Description: "Changes the model if entered",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "new-description",
				Description: "Changes the description if entered",
				Required:    false,
			},
		},
	},

	{
		Name:        "userstats",
		Description: "Check someone's Server Stats",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "Leave blank to see your own stats",
				Required:    false,
			},
		},
	},
	{
		Name:        "ask",
		Description: "Ask Bottom Bot a question, she knows all :3",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "question",
				Description: "Question",
				Required:    true,
			},
		},
	},
	{
		Name:        "meme",
		Description: "Gets a Random Meme",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "search",
				Description: "Optional Search term, leave blank for a Random Subject",
				Required:    false,
			}, {
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "allow-stills",
				Description: "False by default, if set then .jpg/.png's are also included",
				Required:    false,
			},
		},
	}, {
		Name:        "gif-dump",
		Description: "Posts ALL entries in a GIF Bank to a new Thread (Ordered by Newest First)",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "gif-category",
				Description: "The GIF Category to get",
				Required:    true,
			},
		},
	}, {
		Name:        "wow-board",
		Description: "See the server wide WOOOOW leaderboard",
	}, {
		Name:        "roll",
		Description: "Roll Dice!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "d2",
				Description: "How many D2 Dice to roll",
				Required:    false,
			}, {
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "d6",
				Description: "How many D6 Dice to roll",
				Required:    false,
			}, {
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "d10",
				Description: "How many D10 Dice to roll",
				Required:    false,
			}, {
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "d20",
				Description: "How many D20 Dice to roll",
				Required:    false,
			}, {
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "d50",
				Description: "How many D50 Dice to roll",
				Required:    false,
			}, {
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "d100",
				Description: "How many D100 Dice to roll",
				Required:    false,
			},
		},
	},
}

// CommandHandlers ==============================================
var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"fwiday": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		friday.Fwiday(s, i, "w")
		dbhelper.CountCommand("fwiday", i.Member.User.ID)
	},
	"friday": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		friday.Fwiday(s, i, "r")
		dbhelper.CountCommand("fwiday", i.Member.User.ID)
	},
	"reacts": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		animegif.GiveListHandler(s, i)
		dbhelper.CountCommand("reacts", i.Member.User.ID)
	},
	"voicelist": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		fakeyou.GiveTTSHandler(i)
		dbhelper.CountCommand("voicelist", i.Member.User.ID)
	},
	"todolist": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		todo.ToDoList(s, i)
		dbhelper.CountCommand("todolist", i.Member.User.ID)
	},
	"todoadd": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		if !helpers.IsAdmin(i.Member) {
			logging.SendErrorMsgInteraction(i, "Not Allowed", "This is only for Bot Developers!", true)
		} else {
			todo.ToDoAdd(s, i)
			dbhelper.CountCommand("todoadd", i.Member.User.ID)
		}
	},
	"todoupdate": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		if !helpers.IsAdmin(i.Member) {
			logging.SendErrorMsgInteraction(i, "Not Allowed", "This is only for Bot Developers!", true)
		} else {
			todo.ToDoEdit(s, i)
			dbhelper.CountCommand("todoupdate", i.Member.User.ID)
		}
	},
	"tododelete": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		if !helpers.IsAdmin(i.Member) {
			logging.SendErrorMsgInteraction(i, "Not Allowed", "This is only for Bot Developers!", true)
		} else {
			todo.ToDoDelete(s, i)
			dbhelper.CountCommand("tododelete", i.Member.User.ID)
		}
	},
	"ttssearch": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		fakeyou.Search(i)
		dbhelper.CountCommand("ttssearch", i.Member.User.ID)
	},
	"ttsadd": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		if !helpers.IsAdmin(i.Member) {
			logging.SendErrorMsgInteraction(i, "Not Allowed", "This is only for Bot Developers!", true)
		} else {
			fakeyou.Add(i)
			dbhelper.CountCommand("ttsadd", i.Member.User.ID)
		}
	},
	"ttsupdate": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		if !helpers.IsAdmin(i.Member) {
			logging.SendErrorMsgInteraction(i, "Not Allowed", "This is only for Bot Developers!", true)
		} else {
			fakeyou.Update(i)
			dbhelper.CountCommand("ttsupdate", i.Member.User.ID)
		}
	},
	"userstats": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		userstats.GetStats(i)
		dbhelper.CountCommand("userstats", i.Member.User.ID)
	},
	"ask": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		ask.Question(i)
		dbhelper.CountCommand("ask", i.Member.User.ID)
	},
	"meme": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		if i.Member.User.ID == constants.USER_ID_CALLUM {

		}

		meme.GetRandomMeme(i, 0)
		dbhelper.CountCommand("meme", i.Member.User.ID)
	},
	"gif-dump": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		if !helpers.IsAdmin(i.Member) {
			logging.SendErrorMsgInteraction(i, "Not Allowed", "This is only for Bot Developers!", true)
		} else {
			gifbank.DumpCategory(i)
			dbhelper.CountCommand("gif-dump", i.Member.User.ID)
		}
	},
	"slur-definition": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		slurs.DefineASlur(i)
		dbhelper.CountCommand("slur-definition", i.Member.User.ID)
	},
	"wow-board": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		userstats.WowBoard(i)
		dbhelper.CountCommand("wow-board", i.Member.User.ID)
	},
	"roll": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		dice.RollDice(i)
		dbhelper.CountCommand("roll", i.Member.User.ID)
	},
	"bot-help": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if DoNotProcess(i) {
			return
		}
		cmdText := ""
		orderedList := OrderMap(MessageActions)
		for _, list := range orderedList {
			if !list.Value.AdminOnly {
				cmdText += "\n!" + list.Key
			}
		}
		help.GetHelpText(i, cmdText)
		dbhelper.CountCommand("bot-help", i.Member.User.ID)
	},
}

func OrderMap(m map[string]MessageAction) []struct {
	Key   string
	Value MessageAction
} {
	// Extract the keys from the map
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// Sort the keys
	sort.Strings(keys)

	// Create a slice to hold the ordered key-value pairs
	ordered := make([]struct {
		Key   string
		Value MessageAction
	}, len(m))

	// Populate the ordered slice
	for i, key := range keys {
		ordered[i] = struct {
			Key   string
			Value MessageAction
		}{
			Key:   key,
			Value: m[key],
		}
	}

	return ordered
}

func DoNotProcess(i *discordgo.InteractionCreate) bool {
	if config.IsDev {
		if i.ChannelID != constants.CHANNEL_BOT_TEST {
			return true
		}
	} else {
		if i.ChannelID == constants.CHANNEL_BOT_TEST {
			return true
		}
	}
	return false
}
