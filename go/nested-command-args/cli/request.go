package cli

// Request is the internal structure used to pass a command request down the
// handler stack.  Each level in the stack may mutate the Request object.
type Request struct {
	Command string
	Args    []string
}
