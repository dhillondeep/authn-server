package cmd

import (
	"fmt"

	"github.com/keratin/authn-server/app/data"
	"github.com/spf13/cobra"
)

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
		Long:  "Migrate database",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, failed := readEnv()
			if failed {
				return
			}

			fmt.Println("Running migrations.")
			if err := data.MigrateDB(cfg.DatabaseURL); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Migrations complete.")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}
