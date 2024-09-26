package main

import (
	"fmt"
	"os"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/cmd"
)

func main() {
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
