package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"nocturne/internal/app/model"
	"os"
)

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap the storage engine",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := gorm.Open(sqlite.Open("nocturne.db"), &gorm.Config{})

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}

		err = db.AutoMigrate(&model.Snippet{})

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}

		for i := 0; i < 10; i++ {
			db.Create(&model.Snippet{
				SnippetID: uuid.New().String(),
				Title:     faker.New().Lorem().Sentence(5),
				Content:   faker.New().Lorem().Paragraph(5),
			})
		}
		db.Commit()
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bootstrapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bootstrapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
