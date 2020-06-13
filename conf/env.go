package conf

import (
	"net/url"

	"github.com/spf13/viper"
)

type ErrMissingEnvVar string

func (name ErrMissingEnvVar) Error() string {
	return "missing environment variable: " + string(name)
}

func requireEnv(name string) (string, error) {
	if !viper.IsSet(name) {
		return "", ErrMissingEnvVar(name)
	}

	return viper.GetString(name), nil
}

func lookupInt(name string) int {
	return viper.GetInt(name)
}

func lookupBool(name string) bool {
	return viper.GetBool(name)
}

func lookupURL(name string) (*url.URL, error) {
	if !viper.IsSet(name) {
		return nil, nil
	}

	return url.Parse(viper.GetString(name))
}
