package cmd

import (
	"fmt"

	"github.com/keratin/authn-server/conf"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version of authn server",
		Long:  "Version of authn server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\n", conf.Version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
