package app

import (
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var downMigrationCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := setupDB(cmd.Context())
		if err != nil {
			return err
		}

		return goose.Down(db, ".")
	},
}

func init() {
	rootCmd.AddCommand(downMigrationCmd)
}
