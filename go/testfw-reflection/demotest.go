package main

import (
	"log/slog"

	"github.com/egustafson/go/testfw-reflection/demofwk"
)

func init() {
	demofwk.Register(DemoTest)
	demofwk.Register(DemoPanic)
}

func DemoTest(t *demofwk.T) {
	slog.Info("  --> inside fn DemoTest", slog.String("name", t.Name))
}

func DemoPanic(t *demofwk.T) {
	slog.Info("  --> inside fn DemoPanic", slog.String("name", t.Name))
	panic("cause a panic")
	// The framework should catch the panic and mark this test as failed, (and panicked).
}
