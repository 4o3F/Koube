package structs

type KoubeConfig struct {
	Port             string
	EntranceCodeSalt string `mapstructure:"entrance_code_salt"`
	VerifyCodeSalt   string `mapstructure:"verify_code_salt"`
}
