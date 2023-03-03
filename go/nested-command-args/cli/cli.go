package cli

import (
	"context"
	"strings"

	"gopkg.in/yaml.v3"
)

type CliHandler interface {
	Register(handler Handler)
	Execute(ctx context.Context, cmdline string) (Message, error)
}

type Handler struct {
	// Use is the one-line usage message.
	// Example:  add [-F flag-val | -D flag-val] argument-val
	Use string

	// Short
	Short string

	// Command is the fully qualified command string.  If this is a sub-command
	// then the fully qualified string is "parent:sub:command"
	Command string

	// HandlerFn is the function to invoke to handle the command.  When
	// HandlerFn is invoked the command string will be striped from 'args'
	HandlerFn func(ctx context.Context, resp *Response, req *Request)
}

type handlerMap struct {
	Map map[string]Handler
}

var _ CliHandler = (*handlerMap)(nil)

// NewCliHandler returns a new and empty CliHandler.
func NewCliHandler() CliHandler {
	return &handlerMap{
		Map: make(map[string]Handler),
	}
}

// Register registers 'handler' as the handler for 'handler.Command'.  If a
// sub-command is registered then the 'Command' string should contain the full
// sequence of commands to reach this command, colon (:) separated.
func (h *handlerMap) Register(handler Handler) {
	//
	// warning: no error checks
	//
	h.Map[handler.Command] = handler
}

func (h *handlerMap) Execute(ctx context.Context, cmdline string) (Message, error) {

	// parse the command line
	//
	if len(cmdline) < 1 {
		return nil, emptyCommandLineError()
	}
	cl := strings.ReplaceAll(cmdline, "\\\n", " ") // replace escaped newline with space
	lines := strings.Split(cl, "\n")               // only parse to the first "real" newline
	args := strings.Fields(lines[0])
	if len(args) < 1 {
		return nil, emptyCommandLineError() // it's all whitespace
	}

	// locate the handler
	//
	handler, ok := h.Map[args[0]]
	if !ok {
		return nil, unknownCommandError(args[0])
	}

	// process the handler's result --> response + error
	req := &Request{
		Command: args[0],
		Args:    args[1:],
	}
	resp := &Response{}
	handler.HandlerFn(ctx, resp, req)
	message := &message{headers: resp.Headers}
	if resp.Body != nil {
		var contentType string
		var err error
		message.body, contentType, err = transformResultBody(resp.Body)
		if err != nil {
			message.headers.Set(contentTypeKey, contentType) // override because of error
			return message, err                              // return yaml parse error
		}
		if len(contentType) > 0 && !resp.Headers.Contains(contentTypeKey) {
			message.headers.Set(contentTypeKey, contentType)
		}
	}
	return message, resp.Err
}

func transformResultBody(b any) (body string, contentType string, err error) {
	if body, ok := b.(string); ok {
		return body, contentTypeText, nil
	}
	var bodybytes []byte
	bodybytes, err = yaml.Marshal(b)
	if err != nil {
		err := bodyTransformError(err)
		return err.Error(), contentTypeText, err
	}
	return string(bodybytes), contentTypeYaml, nil
}
