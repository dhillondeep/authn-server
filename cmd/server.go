package cmd

import (
	"fmt"
	"os"

	"github.com/keratin/authn-server/conf"
	"github.com/keratin/authn-server/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run authn server",
		Long:  "Run authn server",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, failed := readEnv()
			if failed {
				return
			}

			fmt.Println(fmt.Sprintf("~*~ Keratin AuthN v%s ~*~", conf.Version))

			// Default logger
			logger := logrus.New()
			logger.Formatter = &logrus.JSONFormatter{}
			logger.Level = logrus.DebugLevel
			logger.Out = os.Stdout

			app, err := app.NewApp(cfg, logger)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(fmt.Sprintf("AUTHN_URL: %s", cfg.AuthNURL))
			fmt.Println(fmt.Sprintf("PORT: %d", cfg.ServerPort))
			if app.Config.PublicPort != 0 {
				fmt.Println(fmt.Sprintf("PUBLIC_PORT: %d", app.Config.PublicPort))
			}

			server.Server(app)
		},
	}
)

func init() {
	rootCmd.AddCommand(serverCmd)
}
