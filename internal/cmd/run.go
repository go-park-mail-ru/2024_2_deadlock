package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/depgraph"
)

type RunArgs struct {
	EnvPath string
}

func InitRunCommand() (*cobra.Command, error) {
	args := RunArgs{}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Starts server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			dg := depgraph.NewDepGraph()
			logger, _ := dg.GetLogger()

			viper.SetConfigFile(args.EnvPath)

			var config Config

			err := viper.ReadInConfig()
			if err != nil {
				logger.Warn("Config file Not Found. Using cli arguments")
			} else {
				logger.Debug("Using config file")
				err = viper.Unmarshal(&config)
				if err != nil {
					return err
				}
			}

			logger.Debugw(
				"Got config",
				"args", args,
			)

			return nil
		},
	}

	cmd.Flags().StringVarP(&args.EnvPath, "config", "c", ".env", ".env file path")

	return cmd, nil
}
