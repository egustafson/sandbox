package main

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Handler struct {
	Command string
	Handler func(args []string) (any, error)
}

type HandlerMap map[string]Handler

var RootHandlerMap = make(HandlerMap)

// HandleCommand processes args with the assumption that the first
// argument is a non-flag and is the command.  HandleCommand is
// intended to be used recursively.
func HandleCommand(handlerMap HandlerMap, args []string) (any, error) {
	if len(args) < 1 {
		return nil, errors.New("missing-command")
	}
	handler, ok := handlerMap[args[0]]
	if ok {
		return handler.Handler(args[1:])
	} else {
		return nil, fmt.Errorf("unknown command: %s", args[0])
	}
}

func RootHandler(args []string) (string, error) {
	fmt.Println(strings.Join(args, " "))
	result, err := HandleCommand(RootHandlerMap, args)
	if err != nil {
		return "", err
	}
	out, err := yaml.Marshal(map[string]any{args[0]: result})
	return fmt.Sprintln(string(out)), err
}

func main() {
	fmt.Println("Use go test to exercise this example.")
}
