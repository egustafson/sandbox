package clientlib

// This is the example clientlib package.  It has mock functions that
// use logging from the 'log' package.

import (
	"github.com/rs/zerolog"

	"github.com/egustafson/sandbox/go/zerolog-client/log"
)

func ClientFuncDebug(msg string) {
	dbgEv := log.Logger().Debug()
	dbgEv.Str("caller-msg", msg).Msg("debug output")
	// log.Logger().Debug().Str("caller-msg", msg).Msg("debug output")
}

func ClientFuncInfo(msg string) {
	log.Logger().Info().Str("caller-msg", msg).Msg("info output")
}

func ClientUpLevel() {
	log.SetLevel(zerolog.DebugLevel)
}
