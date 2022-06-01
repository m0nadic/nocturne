package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var host string
var port int
var signingKey string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nocturne",
	Short: "a short composition of a romantic nature, typically for piano",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "host", "b", "localhost", "host address")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 4000, "host port")
	rootCmd.PersistentFlags().StringVarP(&signingKey, "key", "k", "s3cr3t", "signing key")
}
