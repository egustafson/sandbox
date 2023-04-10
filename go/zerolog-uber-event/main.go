package main

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Event zerolog.Event

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	(*Event)(log.Info()).ComplexStruct().Msg("log message") // <- ugly

	LogComplexStruct(log.Info()).Msg("log messsage")

	fmt.Println("done.")
}

// ComplexStruct is an abstracted receiver on zerolog.Event by wrapping it in a
// local class.  This is kinda ugly.
func (ev *Event) ComplexStruct() *zerolog.Event {
	e := (*zerolog.Event)(ev)
	e.Int("key", 42)
	e.Str("mechanism", "receiver")
	return e
}

func LogComplexStruct(e *zerolog.Event) *zerolog.Event {
	e.Int("key", 42)
	e.Str("mechanism", "func")
	return e
}
