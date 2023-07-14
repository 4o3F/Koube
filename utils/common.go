package utils

func InitUtils() {
	initLogger()
	initConfig()

	KoubeLogger.Info().Msg("Init utils complete")
}
