package main

import (
	"log/slog"

	"github.com/egustafson/go/testfw-reflection/demofwk"
)

func init() {
	demofwk.Register(DemoTest)
}

func DemoTest(t *demofwk.T) {
	slog.Info("DemoTest - running", slog.String("name", t.Name))
}
