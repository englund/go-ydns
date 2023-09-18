package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var host string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a YDNS record",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(host)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&host, "host", "H", "", "The host to update")
	updateCmd.MarkFlagRequired("host")
}
