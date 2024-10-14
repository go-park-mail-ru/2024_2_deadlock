package app

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/cmd"
)

var rootCmd = cmd.Init("server")

func MustExecute(ctx context.Context) {
	rootCmd.MustExecute(ctx)
}
