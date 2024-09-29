package app

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/depgraph"
)

type RunArgs struct {
	EnvPath string
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "api",
		Short: "Starts API server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			dg := depgraph.NewDepGraph()
			logger, err := dg.GetLogger()
			if err != nil {
				return errors.Wrap(err, "get logger")
			}

			logger.Info("Starting server...")

			return nil
		},
	})
}
