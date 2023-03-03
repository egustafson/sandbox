package cli

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"
)

type CliHandler interface {
	Register(handler Handler)
	Execute(ctx context.Context, cmdline string) (Message, error)
	DescendHandlers(ctx context.Context, resp *Response, req *Request)
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

// Execute attempts to interpret 'cmdline' and invoke, possibly recursively, the
// correct Handler.HandlerFn.  This method is the external entry point for
// processing (executing) a 'cmdline'.
func (h *handlerMap) Execute(ctx context.Context, cmdline string) (Message, error) {

	// parse the command line
	//
	args, err := translateCmdLineArgs(cmdline)
	if err != nil {
		return nil, err
	}

	// initialize the request and response objects and attempt to execute the HandlerFn
	req := &Request{
		Command:    "",
		Args:       args,
		CmdHandler: h,
	}
	resp := &Response{}
	h.DescendHandlers(ctx, resp, req)

	// process the results of handling the cmdline
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

// DescendHandlers is used to identify and invoke the next Handler in the
// Handler chain.  This method should be used internally to recurse the Handler
// hierarchy.  It may be used by (externally) implemented HandlerFn's if and
// when the Fn needs to recurse further, possibly after processing arguments.
// It is assumed that req.Args[0] is the string identifying the next Handler to
// invoke.  DescendHandlers will invoke the Handler.HandlerFn, or indicate an error.
func (h *handlerMap) DescendHandlers(ctx context.Context, resp *Response, req *Request) {
	for len(req.Args) > 0 {
		// if the first character of req.Args[0] is not an alpha, then its not a
		// command and we've exhausted the search for a command handler.
		if !unicode.IsLetter(rune(req.Args[0][0])) {
			resp.Err = unknownCommandError(req.Command)
			return
		}

		// Consume the next req.Args as the next (sub)command
		if len(req.Command) > 0 {
			req.Command = fmt.Sprintf("%s:%s", req.Command, req.Args[0])
		} else {
			req.Command = req.Args[0] // first element doesn't get a prepended ':'
		}
		req.Args = req.Args[1:]

		// look to see there's a handler for this command sequence
		handler, ok := h.Map[req.Command]
		if ok { // found a handler -> descend
			handler.HandlerFn(ctx, resp, req)
			return // and return, we executed the handler
		}

		// otherwise, attempt to descend further if there's another arg in req.Args
	}
	// if we fell through the loop, there are no more req.Args remaining
	resp.Err = unknownCommandError(req.Command)
}

func transformResultBody(b any) (body_str string, contentType string, err error) {
	if body, ok := b.(string); ok {
		return body, contentTypeText, nil
	}
	var body_bytes []byte
	body_bytes, err = yaml.Marshal(b)
	if err != nil {
		err := bodyTransformError(err)
		return err.Error(), contentTypeText, err
	}
	return string(body_bytes), contentTypeYaml, nil
}

var (
	longFlag_p    = `^\+([A-Za-z].*)$`
	longFlag_re   = regexp.MustCompile(longFlag_p)
	shortFlags_p  = `^\+\+([A-Za-z].*)$`
	shortFlags_re = regexp.MustCompile(shortFlags_p)
)

func translateCmdLineArgs(cmdline string) ([]string, error) {
	if len(cmdline) < 1 {
		return nil, emptyCommandLineError()
	}
	cl := strings.ReplaceAll(cmdline, "\\\n", " ") // replace escaped newline with space
	lines := strings.Split(cl, "\n")               // only parse to the first "real" newline
	if len(lines) > 1 {
		return nil, multiLineCommandLineError(len(lines))
	}

	// extract space separated list of args
	args := strings.Fields(lines[0])
	if len(args) < 1 {
		return nil, emptyCommandLineError() // it's all whitespace
	}

	// translate flag prefixes ('+' -> '--' | '++' -> '-')
	for ii := range args {
		if longFlag_re.MatchString(args[ii]) {
			args[ii] = longFlag_re.ReplaceAllString(args[ii], "--${1}")
		}
		if shortFlags_re.MatchString(args[ii]) {
			args[ii] = shortFlags_re.ReplaceAllString(args[ii], "-${1}")
		}
	}

	return args, nil
}
