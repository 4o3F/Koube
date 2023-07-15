package structs

type KoubeAudience struct {
	Aid          int    `json:"aid"`
	EntranceCode string `json:"entrance_code,omitempty"`
	VerifyCode   string `json:"verify_code,omitempty"`
}
