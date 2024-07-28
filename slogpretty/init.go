package slogpretty

import (
	"log/slog"
	"os"
)

func SetupPrettySlog(level slog.Level) *slog.Logger {
	opts := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}
	
	handler := opts.NewPrettyHandler(os.Stdout)
	
	return slog.New(handler)
}
