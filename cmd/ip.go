package cmd

import (
	"fmt"
	"log"

	"englund.io/ydns/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get IP address",
	Run: func(cmd *cobra.Command, args []string) {
		baseUrl := viper.GetString("baseUrl")
		username := viper.GetString("username")
		password := viper.GetString("password")
		client := pkg.NewYdnsClient(baseUrl, username, password)
		ip, err := client.GetIp()
		if err != nil {
			log.Fatal("error retrieving ip: ", err)
		}
		fmt.Println(*ip)
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
