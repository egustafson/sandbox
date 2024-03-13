package main

import (
	"context"
	"fmt"
)

type ValueKeyType string

const ValueKey ValueKeyType = "MagicValueKey"

func (k ValueKeyType) String() string {
	return string(k)
}

type ValueType int

func (v ValueType) String() string {
	return fmt.Sprintf("%d", v)
}

func InitFn1(ctx context.Context) (context.Context, error) {
	ctx = context.WithValue(ctx, ValueKey, ValueType(1))
	return ctx, nil
}

func InitFn2(ctx context.Context) (context.Context, error) {
	ctx = context.WithValue(ctx, ValueKey, ValueType(2))
	return ctx, nil
}

func HasValue(ctx context.Context, key ValueKeyType) (any, bool) {
	v := ctx.Value(key)
	return v, (v != nil)
}

func main() {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if ctx, err = InitFn1(ctx); err != nil {
		panic(err)
	}
	if ctx, err = InitFn2(ctx); err != nil {
		panic(err)
	}
	if v, ok := HasValue(ctx, ValueKey); ok {
		fmt.Printf("ctx['%s']: %v\n", ValueKey, v)
	} else {
		fmt.Printf("ctx does NOT have '%s'\n", ValueKey)
	}
	fmt.Printf("ctx: %v\n", ctx)
}
