package database

import (
	"database/sql"
	"errors"
	"github.com/4o3F/Koube/utils"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

var KoubeDatabase *sql.DB

func InitDatabase() {
	database, err := sql.Open("sqlite3", "file:database.db?_journal=WAL&chache=shared")
	database.SetMaxOpenConns(1)
	KoubeDatabase = database
	if err != nil {
		utils.KoubeLogger.Fatal().Err(err).Send()
	}
	utils.KoubeLogger.Info().Msg("Database init complete")

	if !utils.KoubeConfig.GenerationComplete {
		err := GenerateAudiencesData()
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				if strings.Contains(err.Error(), "verify_code") {
					utils.KoubeLogger.Fatal().Msg("Verify code collision detected, please increase the length of verify code")
				} else if strings.Contains(err.Error(), "entrance_code") {
					utils.KoubeLogger.Fatal().Msg("Entrance code collision detected, please increase the length of entrance code")
				} else {
					utils.KoubeLogger.Fatal().Err(err).Send()
				}
			}
			utils.KoubeLogger.Fatal().Err(err).Send()
		} else if err != nil {
			utils.KoubeLogger.Fatal().Err(err).Send()
		}
		err = utils.SwitchAudienceGenerationStatus()
		if err != nil {
			utils.KoubeLogger.Fatal().Err(err).Send()
		}
		utils.KoubeLogger.Info().Msg("Audiences data generation complete")
	} else {
		utils.KoubeLogger.Info().Msg("Audiences data already generated, skipping")
	}
}
