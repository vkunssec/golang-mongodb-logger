package core

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

func Logger() *zerolog.Logger {
	buildInfo, _ := debug.ReadBuildInfo()

	logger := zerolog.
		New(os.Stderr).
		Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
			FormatCaller: func(i interface{}) string {
				return filepath.Base(fmt.Sprintf("%s", i))
			},
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %s |", i)
			},
			TimeFormat: time.RFC822,
		}).
		With().
		Caller().
		Timestamp().
		Str("go_version", buildInfo.GoVersion).
		Logger()

	return &logger
}
