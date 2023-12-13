package cmd

import (
	"log"

	"englund.io/ydns/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ydns",
	Short: "A YDNS command-line tool.",
}

func Execute() {
	config.InitConfig()
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("something went terrible wrong: %s", err)
	}
}
