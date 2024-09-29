package app

import (
	"github.com/spf13/cobra"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/app"
)

type RunArgs struct {
	EnvPath string
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "api",
		Short: "Starts API server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return app.Run(cmd.Context(), &app.APIEntrypoint{Config: rootCmd.Config})
		},
	})
}
