package main

import (
	"fmt"
	"os"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/cmd"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/depgraph"
)

func main() {
	dg := depgraph.NewDepGraph()
	logger, err := dg.GetLogger()
	if err != nil {
		fmt.Println("get logger: %w", err)
		os.Exit(1)
	}

	logger.Info("Starting server...")

	rootCmd, err := cmd.InitRunCommand()
	if err != nil {
		fmt.Println("init run command: %w", err)
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("execute: %w", err)
		os.Exit(1)
	}
}