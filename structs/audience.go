package structs

type KoubeAudience struct {
	Aid          int    `json:"aid"`
	EntranceCode string `json:"entrance_code"`
	VerifyCode   string `json:"verify_code"`
}
