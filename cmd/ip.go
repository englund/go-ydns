package cmd

import (
	"fmt"
	"log"

	"englund.io/ydns/pkg"
	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get IP address",
	Run: func(cmd *cobra.Command, args []string) {
		client := pkg.NewYdnsClient(cfg.BaseUrl, cfg.Username, cfg.Password)
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
