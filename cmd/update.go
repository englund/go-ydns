package cmd

import (
	"log"

	"englund.io/ydns/pkg"
	"github.com/spf13/cobra"
)

var hosts []string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update one or more YDNS records",
	Run: func(cmd *cobra.Command, args []string) {
		client := pkg.NewYdnsClient(cfg.BaseUrl, cfg.Username, cfg.Password)
		ip, err := client.GetIp()
		if err != nil {
			log.Fatal("error retrieving ip: ", err)
		}

		for _, host := range hosts {
			if err := client.Update(host, *ip); err != nil {
				log.Fatalf("error updating host \"%s\": %s", host, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringSliceVarP(&hosts, "host", "H", nil, "One or more hosts to update")
	updateCmd.MarkFlagRequired("host")
}
