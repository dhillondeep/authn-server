package cmd

import (
	"fmt"
	"os"

	"github.com/keratin/authn-server/conf"
)

// readEnv reads configuration from environment values
func readEnv() (*conf.Config, bool) {
	cfg, err := conf.ReadEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		fmt.Fprintf(os.Stderr, "%s\n", "see: https://github.com/keratin/authn-server/blob/master/docs/config.md")
		return cfg, true
	}

	return cfg, false
}
