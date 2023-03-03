package cli

const (
	contentTypeKey  = "Content-Type"
	contentTypeText = "text/plain"
	contentTypeYaml = "text/yaml"
)

// Header is a single key/value header field where both the key and the value
// are strings in the spirit of RFC5322, although necessarily conformant.
type Header struct {
	Key   string
	Value string
}

// Headers is aliased so that methods can be implemented
type Headers []Header

// --  Headers implementation  ----------------------------

// Keys returns a list of unique keys contained in the Headers
func (h *Headers) Keys() []string {
	keySet := make(map[string]struct{})
	for _, hdr := range *h {
		keySet[hdr.Key] = struct{}{}
	}
	keys := make([]string, 0, len(keySet))
	for k := range keySet {
		keys = append(keys, k)
	}
	return keys
}

// Contains returns true if 'k' is a key in h
func (h *Headers) Contains(k string) bool {
	for _, hdr := range *h {
		if hdr.Key == k {
			return true
		}
	}
	return false
}

// Get returns the value of the first header with key 'k' or the zero value
// string, ("")
func (h *Headers) Get(k string) string {
	for _, hdr := range *h {
		if hdr.Key == k {
			return hdr.Value
		}
	}
	// key not found return zero value
	return ""
}

// GetAll returns a list of string for each header with key 'k'
func (h *Headers) GetAll(k string) []string {
	values := make([]string, 0)
	for _, hdr := range *h {
		if hdr.Key == k {
			values = append(values, hdr.Value)
		}
	}
	return values
}

// Set sets the first header with key = 'k' to 'v' or appends a new Header to
// the list.
func (h *Headers) Set(k, v string) {
	for idx, hdr := range *h {
		if hdr.Key == k {
			(*h)[idx].Value = v // overwrite
			return
		}
	}
	*h = append(*h, Header{Key: k, Value: v})
}

func (h *Headers) Append(k, v string) {
	*h = append(*h, Header{Key: k, Value: v})
}

// Delete removes the first header with key 'k'
func (h *Headers) Delete(k string) {
	for i, hdr := range *h {
		if hdr.Key == k {
			// non-order preserving
			(*h)[i] = (*h)[len(*h)-1]
			*h = (*h)[:len(*h)-1]
			return
		}
	}
}

// DeleteAll removes all headers with key 'k'
func (h *Headers) DeleteAll(k string) {
	// iterate from last to first (i.e. reverse)
	for i := len(*h) - 1; i > -1; i-- {
		if (*h)[i].Key == k {
			// non-order preserving
			(*h)[i] = (*h)[len(*h)-1]
			*h = (*h)[:len(*h)-1]
		}
	}
}
