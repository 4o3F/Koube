package structs

type KoubeConfig struct {
	Port                  string
	EntranceCodeSalt      string `mapstructure:"entrance_code_salt"`
	MaxEntranceCodeLength int    `mapstructure:"max_entrance_code_length"`
	VerifyCodeSalt        string `mapstructure:"verify_code_salt"`
	MaxVerifyCodeLength   int    `mapstructure:"max_verify_code_length"`
	APIAuthKey            string `mapstructure:"api_auth_key"`
	MaxAudienceCount      int    `mapstructure:"max_audience_count"`
	GenerationComplete    bool   `mapstructure:"generation_complete"`
	ShowStartTime         int64  `mapstructure:"show_start_time"`
	AccessAllowDuration   int64  `mapstructure:"access_allow_duration"`
}
