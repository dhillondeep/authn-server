package conf

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("BCRYPT_COST", 11)
	viper.SetDefault("PASSWORD_POLICY_SCORE", 2)
	viper.SetDefault("REFRESH_TOKEN_TTL", 86400*30)
	viper.SetDefault("PASSWORD_RESET_TOKEN_TTL", 1800)
	viper.SetDefault("PASSWORDLESS_TOKEN_TTL", 1800)
	viper.SetDefault("ACCESS_TOKEN_TTL", 3600)
	viper.SetDefault("DAILY_ACTIVES_RETENTION", 365)
	viper.SetDefault("WEEKLY_ACTIVES_RETENTION", 104)
	viper.SetDefault("PUBLIC_PORT", 0)

	viper.SetDefault("PASSWORD_CHANGE_LOGOUT", false)
	viper.SetDefault("USERNAME_IS_EMAIL", false)
	viper.SetDefault("ENABLE_SIGNUP", true)
	viper.SetDefault("PROXIED", false)
}
