package cli

// Response is the internal structure used to return results from a command.  A
// Response may be modified by each nested command as it is returned up the stack.
type Response struct {
	Headers Headers
	Body    any
	Err     error
}
