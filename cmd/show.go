package cmd

import (
	"fmt"
	"nocturne/internal/pkg/cleint/snippet"
	"os"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays the details of a snippet",
	Run: func(cmd *cobra.Command, args []string) {
		client := snippet.NewClient(host, port, signingKey)

		for _, arg := range args {

			snippet, err := client.GetSnippet(arg)

			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
				os.Exit(1)
			}

			fmt.Println("ID:", snippet.SnippetID)
			fmt.Println("Title:", snippet.Title)
			fmt.Println("Content:", snippet.Content)

		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
