package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/keratin/authn-server/conf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	rootCmd    = &cobra.Command{
		Use:     "authn-server",
		Version: conf.Version,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file for authn-server")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func initConfig() {
	viper.SetTypeByDefaultValue(true)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if configFile != "" {
		viper.SetConfigFile(configFile)

		if err := viper.ReadInConfig(); err != nil {
			fmt.Fprintf(os.Stderr, "Could not read config file %s: %s\n", configFile, err.Error())
			os.Exit(1)
		}
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			// this file is optional and hence only other errors matter
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				fmt.Fprintf(os.Stderr, "Could not parse .env file: %s\n", err.Error())
			}
		}
	}
}
