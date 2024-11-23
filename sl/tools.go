package sl

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

func Err(err error) slog.Attr {
	if err == nil {
		err = fmt.Errorf("This error is nil")
	}
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func Fatal(logger *slog.Logger, err error) {
	logger.Error("Fatal error:", Err(err))
	os.Exit(1)
}

func Panic(logger *slog.Logger, err error) {
	log.Panic("Panic error:" + err.Error())
}
