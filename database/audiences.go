package database

import (
	"database/sql"
	"github.com/4o3F/Koube/structs"
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

	// CREATE NEW CLEAN TABLE
	stmt, err = KoubeDatabase.Prepare("CREATE TABLE `audience` (`aid` int UNIQUE PRIMARY KEY,`entrance_code` varchar(255) NOT NULL,`verify_code` varchar(255) NOT NULL UNIQUE,`entered` boolean NOT NULL DEFAULT false)")
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

func GetAudienceByAid(aid int) (structs.KoubeAudience, error) {
	stmt, err := KoubeDatabase.Prepare("SELECT `entrance_code`,`verify_code` FROM `audience` WHERE `aid` = ?")
	if err != nil {
		return structs.KoubeAudience{}, err
	}
	rows, err := stmt.Query(aid)
	if err != nil {
		return structs.KoubeAudience{}, err
	}
	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if err == nil {
			err = closeErr
		}
	}(rows)
	var audience structs.KoubeAudience
	for rows.Next() {
		err := rows.Scan(&audience.EntranceCode, &audience.VerifyCode)
		if err != nil {
			return structs.KoubeAudience{}, err
		}
	}
	return audience, nil
}

func AudienceEnter(aid int) error {
	stmt, err := KoubeDatabase.Prepare("UPDATE `audience` SET `entered` = true WHERE `aid` = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(aid)
	if err != nil {
		return err
	}
	return nil
}
