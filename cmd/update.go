package cmd

import (
	"englund.io/ydns/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var host string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a YDNS record",
	Run: func(cmd *cobra.Command, args []string) {
		username := viper.GetString("username")
		password := viper.GetString("password")
		client := pkg.NewYdnsClient(&username, &password)
		client.Update(&host)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&host, "host", "H", "", "The host to update")
	updateCmd.MarkFlagRequired("host")
}
