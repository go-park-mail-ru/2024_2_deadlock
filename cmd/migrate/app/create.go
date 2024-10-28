package app

import (
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var createMigrationCmd = &cobra.Command{
	Use:   "create",
	Short: "create new migration file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return goose.Create(nil, "migrations", args[0], "sql")
	},
}

func init() {
	rootCmd.AddCommand(createMigrationCmd)
}
