package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nocturne/internal/pkg/cleint/snippet"
	"os"
)

var (
	title   string
	content string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a snippet in server.",
	Run: func(cmd *cobra.Command, args []string) {
		client := snippet.NewClient(host, port, signingKey)
		snippetID, err := client.CreateSnippet(title, content)

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}

		fmt.Println("Created Snippet:", snippetID)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(&title, "title", "", "title of the snippet")
	createCmd.Flags().StringVar(&content, "content", "", "content of the snippet")
}
