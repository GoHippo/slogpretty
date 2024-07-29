package slogpretty

import (
	"fmt"
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

func SetupPrettySlogInFile(level slog.Level, path_file string) *slog.Logger {
	opts := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}
	
	file, err := os.OpenFile(path_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file for logger:", err)
		os.Exit(1)
	}
	
	handler := opts.NewPrettyHandler(file)
	
	return slog.New(handler)
}
