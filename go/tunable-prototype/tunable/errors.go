package tunable

import "fmt"

type NotExistError interface {
	error
}

func notExistError(name string) NotExistError {
	e := fmt.Errorf("tunable does not exist: %s", name)
	return NotExistError(e)
}

type TypeMismatchError interface {
	error
}

func typeMismatchError(received any, expected string) TypeMismatchError {
	e := fmt.Errorf("type mismatch: %T != %s", received, expected)
	return TypeMismatchError(e)
}

type ParseError interface {
	error
}

func parseError(err error) ParseError {
	e := fmt.Errorf("tunable parse error: %w", err)
	return ParseError(e)
}

func fatalError(err error) {
	panic(err)
}
