package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

var Config config

type config struct {
	BaseUrl    string `mapstructure:"baseUrl"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	LastIpFile string `mapstructure:"lastIpFile"`
}

func (c *config) validate() error {
	if c.BaseUrl == "" {
		return errors.New("baseUrl is required")
	}
	if c.Username == "" {
		return errors.New("username is required")
	}
	if c.Password == "" {
		return errors.New("password is required")
	}
	if c.LastIpFile == "" {
		return errors.New("lastIpFile is required")
	}
	return nil
}

func InitConfig() {
	viper.SetConfigName("ydns")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/ydns-updater/")
	viper.AddConfigPath("$HOME/.ydns-updater")
	viper.AddConfigPath(".")

	viper.SetDefault("lastIpFile", "/tmp/ydns_last_ip")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	err = Config.validate()
	if err != nil {
		log.Fatalf("invalid configuration: %s", err)
	}
}
