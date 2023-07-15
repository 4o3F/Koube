package database

import (
	"github.com/4o3F/Koube/utils"
	"strconv"
)

func GenerateAudiencesData() error {
	// DROP EXISTING TABLE
	stmt, err := KoubeDatabase.Prepare("DROP TABLE `audience`")
	if err != nil {
		return err
	}
	_, _ = stmt.Exec()

	stmt, err = KoubeDatabase.Prepare("CREATE TABLE `audience` (`aid` int UNIQUE PRIMARY KEY,`entrance_code` varchar(255) NOT NULL,`verify_code` varchar(255) NOT NULL UNIQUE)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	stmt, err = KoubeDatabase.Prepare("INSERT INTO `audience` (`aid`,`entrance_code`,`verify_code`) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	for i := 0; i < utils.KoubeConfig.MaxAudienceCount; i++ {
		entranceCode := utils.GenerateHash(strconv.Itoa(i) + utils.KoubeConfig.EntranceCodeSalt)[:utils.KoubeConfig.MaxEntranceCodeLength]
		verifyCode := utils.GenerateHash(strconv.Itoa(i) + utils.KoubeConfig.VerifyCodeSalt)[:utils.KoubeConfig.MaxVerifyCodeLength]
		utils.KoubeLogger.Info().Msgf("\u001B[1A\u001B[KGenerating audience data: " + strconv.Itoa(i) + " " + entranceCode + " " + verifyCode)
		_, err = stmt.Exec(i, entranceCode, verifyCode)
		if err != nil {
			return err
		}
	}
	return nil
}
