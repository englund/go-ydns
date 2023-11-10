package cmd

import (
	"fmt"
	"log"

	"englund.io/ydns/pkg"
	"github.com/spf13/cobra"
)

var hosts []string

func getIp(client *pkg.YdnsClient) string {
	ip, err := client.GetIp()
	if err != nil {
		log.Fatal("error retrieving ip: ", err)
	}
	return *ip
}

func updateHost(client *pkg.YdnsClient, host string, ip string) {
	if err := client.Update(host, ip); err != nil {
		log.Fatalf("error updating host \"%s\": %s", host, err)
	}
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update one or more YDNS records",
	Run: func(cmd *cobra.Command, args []string) {
		client := pkg.NewYdnsClient(cfg.BaseUrl, cfg.Username, cfg.Password)
		currentIp := getIp(client)

		savedIp, err := pkg.ReadIpFromFile(cfg.LastIpFile)
		if err != nil {
			log.Fatal(err)
		}

		if currentIp == savedIp {
			// IP address has not changed, no need to update
			return
		}

		if err := pkg.WriteIpToFile(cfg.LastIpFile, currentIp); err != nil {
			log.Fatal(err)
		}

		for _, host := range hosts {
			updateHost(client, host, currentIp)
		}

		fmt.Printf("Successfully updated hosts with new IP address: %s\n", currentIp)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringSliceVarP(&hosts, "host", "H", nil, "One or more hosts to update")
	updateCmd.MarkFlagRequired("host")
}
