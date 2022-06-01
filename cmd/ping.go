/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"nocturne/internal/pkg/cleint/snippet"
	"os"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Pings the server",
	Run: func(cmd *cobra.Command, args []string) {
		client := snippet.NewClient(host, port, signingKey)
		ping, err := client.Ping()

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}

		fmt.Println(ping)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
