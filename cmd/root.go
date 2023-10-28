package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ydns",
	Short: "A YDNS command-line tool.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("something went terrible wrong: %s", err)
	}
}

func init() {
	viper.SetConfigName("ydns")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/ydns/")
	viper.AddConfigPath("$HOME/.ydns")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
}
