package cli

import (
	"errors"
	"fmt"
)

type EmptyCommandLineError interface {
	error
}

func emptyCommandLineError() EmptyCommandLineError {
	e := errors.New("empty command line")
	return EmptyCommandLineError(e)
}

type UnknownCommandError interface {
	error
}

func unknownCommandError(name string) UnknownCommandError {
	e := fmt.Errorf("unknown command: %s", name)
	return UnknownCommandError(e)
}

type BodyTransformError interface {
	error
}

func bodyTransformError(er error) BodyTransformError {
	e := fmt.Errorf("body transformation failed: %w", er)
	return BodyTransformError(e)
}
