package log

// This package implemts a prototype log package ON TOP of zerolog.
// It's goal is to provide a zerolog logger for a client library that
// is insulated from main application manipulation of zerolog and its
// globals.

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

var (
	logOut zerolog.ConsoleWriter
	logger zerolog.Logger

	defaultLevel   = zerolog.InfoLevel
	logLevel       = defaultLevel // override in applyEnvVars()
	levelOverriden = false
	globalOverride = false
	warningLogged  = false
)

func init() {
	applyEnvVars() // read and apply env vars

	logOut = zerolog.NewConsoleWriter()
	logger = zerolog.New(logOut).Level(logLevel)
}

// Logger returns the client lib logger and performs some checks to
// make sure client lib logging isn't inadvertantly masked by the
// zerolog.GlobalLevel(), and mitigates if it is.
func Logger() *zerolog.Logger {
	if logLevel < zerolog.GlobalLevel() {
		fmt.Println("--> logLevel < globalLevel")
		if globalOverride {
			zerolog.SetGlobalLevel(logLevel)
			logger.Warn().Str("new-level", logLevel.String()).Msg("zerolog global log level overriden")
		} else if levelOverriden && !warningLogged {
			logger.Warn().Str("sandbox-log-level", logLevel.String()).Msg("zerolog global level masks log messages")
			warningLogged = true
		}
	}
	return &logger
}

// SetLevel allows programatic setting of the client libraries desired
// log level.
func SetLevel(lvl zerolog.Level) {
	fmt.Printf("--> log level (re)set: %s\n", lvl.String())
	levelOverriden = true
	logLevel = lvl
	logger = zerolog.New(logOut).Level(lvl)
}

// SetGlobalOverride allows programatic setting of the globalOverride
// flag, which allows the zerolog.GlobalLevel() to be reset to a more
// verbose logging level.
func SetGlobalOverride(f bool) {
	globalOverride = f
	if f {
		fmt.Println("--> global override set TRUE by program")
	} else {
		fmt.Println("--> global override set FALSE by program")
	}
}

// applyEnvVars parses environment variables and uses them to set the
// client logging libraries configuration.
func applyEnvVars() {

	// env-var to set the log level
	//
	envLvl := os.Getenv("SANDBOX_LOG_LEVEL")
	if len(envLvl) > 0 {
		envLvl = strings.ToLower(envLvl)
		level, err := zerolog.ParseLevel(envLvl)
		if err == nil {
			logLevel = level
			levelOverriden = true
			fmt.Printf("--> log level override by env: %s\n", envLvl)
		}
	}

	// env-var to override (reset) the zerolog.GlobalLevel()
	//
	envGlobalOverride := os.Getenv("SANDBOX_LEVEL_OVERRIDE")
	if len(envGlobalOverride) > 0 {
		if strings.ToUpper(envGlobalOverride) == "TRUE" {
			globalOverride = true
			fmt.Println("--> global override set TRUE by envvar")
		} else {
			globalOverride = false
			fmt.Println("--> global override set FALSE by envvar")
		}
	}
}
