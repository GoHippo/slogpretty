package STDLogger

import (
	"fmt"
	"os"
)

func New(path string, terminal_write bool) *STDLogger {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening STDLogger for logger:", err)
		os.Exit(1)
	}

	var terminal *os.File
	if terminal_write {
		terminal = os.Stdout
	}

	return &STDLogger{
		file:     file,
		terminal: terminal,
	}
}

type STDLogger struct {
	file     *os.File
	terminal *os.File
}

func (s STDLogger) Write(p []byte) (n int, err error) {
	if s.terminal != nil {
		n, err = os.Stdout.Write(p)
		if err != nil {
			return n, err
		}
	}

	return s.file.Write(p)
}
