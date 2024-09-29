package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
)

const (
	cmdConfigName      = "config"
	cmdConfigShorthand = "c"
	cmdConfigValue     = "dev.yaml"
	cmdConfigUsage     = ".yaml file path"
)

type Options struct {
	ConfigPath string
}

type RootCommand struct {
	*cobra.Command
	Config  *bootstrap.Config
	Options Options
}

func Init(name string) *RootCommand {
	cmd := &RootCommand{
		Command: &cobra.Command{
			Use: name,
		},
	}
	cobra.OnInitialize(cmd.setup)

	cmd.Flags().StringVarP(
		&cmd.Options.ConfigPath,
		cmdConfigName,
		cmdConfigShorthand,
		cmdConfigValue,
		cmdConfigUsage,
	)

	return cmd
}

func (c *RootCommand) Execute(ctx context.Context) error {
	return c.ExecuteContext(ctx)
}

func (c *RootCommand) MustExecute(ctx context.Context) {
	if err := c.Execute(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "app failed: %v\n", err)
		os.Exit(1)
	}
}

func (c *RootCommand) setup() {
	cfg, err := bootstrap.Setup(c.Options.ConfigPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "setup failed: %v\n", err)
		os.Exit(2)
	}

	c.Config = cfg
}
