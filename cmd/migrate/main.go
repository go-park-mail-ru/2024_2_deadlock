package main

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/cmd/migrate/app"
)

func main() {
	app.MustExecute(context.Background())
}
