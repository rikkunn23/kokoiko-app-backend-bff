package wslog

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

var logger *slog.Logger

// New ..
func New(addSource bool, level slog.Level) {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   addSource,
		Level:       level,
		ReplaceAttr: replaceSourceAttr,
	}))
}

// Get ..
func Get() *slog.Logger {
	return logger
}

func replaceSourceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Key == "source" {
		_, file, line, ok := runtime.Caller(7)
		if ok {
			return slog.String(a.Key, fmt.Sprintf("%s:%d", file, line))
		}
	}
	return a
}
