package logger

import (
	"fmt"
	"time"

	config "github.com/dabi-ngin/discgo-bot/Config"
)

// Logs important Events, ie. The requesting/delivery of user input
func Event(guildId string, logText string, a ...any) {
	if config.LoggingLevel <= config.LoggingLevelEvent {
		infoLine := fmt.Sprintf("%v | %v", time.Now().Format("02/01/06 15:04:05.000"), GetStack())
		if guildId != "" {
			infoLine += " | " + guildId
		}
		formattedLogText := logText
		if len(a) > 0 {
			formattedLogText = FormatInboundLogText(logText, a...)
		}
		SendToConsole(infoLine, formattedLogText, config.LoggingLevelEvent)
		if config.IsDev {
			infoLine += " | " + config.HostName
		}
		SendLogToDiscord(infoLine, formattedLogText, config.LoggingLevelEvent)
	}
}
