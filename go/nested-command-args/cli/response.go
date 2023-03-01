package cli

// Response is the returned result of passing a command string into the
// CliHandler's Execute() function.
type Response interface {
	Headers() []Header
	Body() string
}

// Result is the internal structure used to return results from a command.  A
// Result may be modified by each nested command as it is returned up the stack.
type Result struct {
	Headers Headers
	Body    any
	Err     error
}

type response struct {
	headers Headers
	body    string
}

// static check: struct response implements Response
var _ Response = (*response)(nil)

func (r *response) Headers() []Header {
	return r.headers
}

func (r *response) Body() string {
	return r.body
}
