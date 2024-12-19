package slogpretty

import (
	"log/slog"
	"testing"
)

func TestSTDLogger(t *testing.T) {
	Log := SetupPrettySlogInFile(slog.LevelInfo, "log.txt", true)
	Log.Info("Write in file & terminal start!")
	Log.Error("Error test in file & terminal")

	Log2 := SetupPrettySlogInFile(slog.LevelInfo, "log.txt", false)
	Log2.Info("Write in file only start!")
	Log2.Error("Error test in file")
}
