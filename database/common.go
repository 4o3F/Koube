package database

import (
	"github.com/4o3F/Koube/utils"
	"github.com/nutsdb/nutsdb"
)

var KoubeDatabase *nutsdb.DB

func InitDatabase() {
	database, err := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir("./database"),
	)
	KoubeDatabase = database
	if err != nil {
		utils.KoubeLogger.Fatal().Err(err).Send()
	}
	utils.KoubeLogger.Info().Msg("Database init complete")
}
