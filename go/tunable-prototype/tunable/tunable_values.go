package tunable

import "strconv"

// This package is heavily inspired by the golang flag and spf13/pflags packages.

type Value interface {
	String() string
	Set(s string) error
	Get() any
}

// --  Boolean Value  -----------------

type boolValue bool

func newBoolValue(v bool, p *bool) *boolValue {
	*p = v
	return (*boolValue)(p)
}

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return err // TODO: wrap error
	}
	*b = boolValue(v)
	return nil
}

func (b *boolValue) Get() any {
	return bool(*b)
}

func (b *boolValue) String() string {
	return strconv.FormatBool(bool(*b))
}

// --  Integer Value  -----------------

type intValue int

func newIntValue(v int, p *int) *intValue {
	*p = v
	return (*intValue)(p)
}

func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err // TODO: wrap error
	}
	*i = intValue(v)
	return nil
}

func (i *intValue) Get() any {
	return int(*i)
}

func (i *intValue) String() string {
	return strconv.FormatInt(int64(*i), 10)
}
