package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nocturne/internal/pkg/cleint/snippet"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list snippets",
	Run: func(cmd *cobra.Command, args []string) {
		client := snippet.NewClient(host, port, signingKey)
		snippets, err := client.GetSnippets()

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}

		for _, snip := range snippets {
			fmt.Println(snip.SnippetID, snip.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
