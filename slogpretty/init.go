package slogpretty

import (
	"github.com/GoHippo/slogpretty/slogpretty/STDLogger"
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

func SetupPrettySlogInFile(level slog.Level, path_file string, terminal_write bool) *slog.Logger {
	opts := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}

	handler := opts.NewPrettyHandler(STDLogger.New(path_file, terminal_write))

	return slog.New(handler)
}
