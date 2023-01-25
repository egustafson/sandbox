package tunable

import (
	"strconv"
)

// This package is heavily inspired by the golang flag and spf13/pflags packages.

type Value interface {
	String() string
	Set(s string) error
	SetValue(any) error
	Get() any
	Clone() Value
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
		return parseError(err)
	}
	*b = boolValue(v)
	return nil
}

func (b *boolValue) SetValue(v any) error {
	if str, ok := v.(string); ok {
		return b.Set(str)
	}
	if val, ok := v.(bool); ok {
		*b = boolValue(val)
		return nil
	}
	return typeMismatchError(v, "bool")
}

func (b *boolValue) Get() any {
	return bool(*b)
}

func (b *boolValue) String() string {
	return strconv.FormatBool(bool(*b))
}

func (b *boolValue) Clone() Value {
	clone := new(bool)
	*clone = bool(*b)
	return (*boolValue)(clone)
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
		return parseError(err)
	}
	*i = intValue(v)
	return nil
}

func (i *intValue) SetValue(v any) error {
	if str, ok := v.(string); ok {
		return i.Set(str)
	}
	if val, ok := v.(int); ok {
		*i = intValue(val)
		return nil
	}
	return typeMismatchError(v, "int")
}

func (i *intValue) Get() any {
	return int(*i)
}

func (i *intValue) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

func (i *intValue) Clone() Value {
	clone := new(int)
	*clone = int(*i)
	return (*intValue)(clone)
}
