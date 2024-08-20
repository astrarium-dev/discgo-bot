package database

import (
	cache "github.com/dabi-ngin/discgo-bot/Cache"
	logger "github.com/dabi-ngin/discgo-bot/Logger"
)

func LogCommandUsage(guildId string, userId string, commandTypeId int, command string) {

	// 1. Try an Update
	updateQuery := `
		UPDATE CommandLog
		SET Count = Count+1, LastUsedDateTime = NOW()
		WHERE GuildID = ? AND UserID = ? AND CommandTypeID = ? AND Command = ?
	`

	result, err := db.Exec(updateQuery, guildId, userId, commandTypeId, command)
	if err != nil {
		logger.Error(guildId, err)
		return
	}

	// 2. Check if we affected a row
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error(guildId, err)
		return
	}

	if rowsAffected == 0 {
		// 3. If not, insert the new row
		insertQuery := `
			INSERT INTO CommandLog
			(GuildID, UserID, CommandTypeID, Command)
			VALUES
			(?, ?, ?, ?)
		`

		_, err = db.Exec(insertQuery, guildId, userId, commandTypeId, command)
		if err != nil {
			logger.Error(guildId, err)
			return
		}
	}

	// 4. Update the Guild Cache
	cache.UpdateGuildLastCommand(guildId)
}
