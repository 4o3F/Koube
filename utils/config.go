package utils

import (
	"github.com/4o3F/Koube/structs"
	"github.com/spf13/viper"
)

var KoubeConfig structs.KoubeConfig

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		KoubeLogger.Fatal().Err(err).Send()
	}
	err = viper.Unmarshal(&KoubeConfig)
	if err != nil {
		KoubeLogger.Fatal().Err(err).Send()
	}
	KoubeLogger.Info().Msg("Init config complete")
}

func SwitchAudienceGenerationStatus() error {
	KoubeConfig.GenerationComplete = !KoubeConfig.GenerationComplete
	viper.Set("generation_complete", KoubeConfig.GenerationComplete)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
