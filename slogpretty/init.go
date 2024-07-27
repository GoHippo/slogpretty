package slogpretty

import (
	"log/slog"
	"os"
)

func SetupPrettySlog() *slog.Logger {
	opts := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	
	handler := opts.NewPrettyHandler(os.Stdout)
	
	return slog.New(handler)
}
