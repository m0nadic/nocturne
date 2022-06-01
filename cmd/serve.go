package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"nocturne/internal/app/server"
	"os"
)

var dbPath string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Starting server at %s:%d with %s", host, port, dbPath)
		err := server.InitHttpServer(host, port, dbPath)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&dbPath, "db", "d", "./nocturne.db", "full path of sqlite db")
}
