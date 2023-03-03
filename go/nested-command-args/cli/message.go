package cli

// Message is the returned result of passing a command string into the
// CliHandler's Execute() function.
type Message interface {
	Headers() []Header
	Body() string
}

type message struct {
	headers Headers
	body    string
}

// static check: struct response implements Response
var _ Message = (*message)(nil)

func (r *message) Headers() []Header {
	return r.headers
}

func (r *message) Body() string {
	return r.body
}
