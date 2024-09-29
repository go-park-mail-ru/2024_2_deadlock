package app

import (
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var upMigrationCmd = &cobra.Command{
	Use:   "up",
	Short: "run up migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := setupDB(cmd.Context())
		if err != nil {
			return err
		}

		return goose.Up(db, ".")
	},
}

func init() {
	rootCmd.AddCommand(upMigrationCmd)
}
