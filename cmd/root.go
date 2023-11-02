package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	BaseUrl  string `mapstructure:"baseUrl"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func (c *Config) Validate() error {
	if c.BaseUrl == "" {
		return errors.New("baseUrl is required")
	}
	if c.Username == "" {
		return errors.New("username is required")
	}
	if c.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

var cfg Config

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

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	err = cfg.Validate()
	if err != nil {
		log.Fatalf("invalid configuration: %s", err)
	}
}
