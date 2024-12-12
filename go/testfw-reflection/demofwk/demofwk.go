package demofwk

import "context"

type T struct {
	PublicField string
	hiddenField string

	ctx context.Context
}
