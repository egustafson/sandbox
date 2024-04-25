package tunables

import (
	"fmt"
	"strconv"
)

type Value interface {
	Set(s string) (Value, error)
	String() string
}

type Tunable interface {
	Name() string
	Value() Value
	Override(v Value) error
	Reset()
	IsOverriden() bool
}

type IntegerVal int64

func (i IntegerVal) Set(s string) (Value, error) {
	j, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	i = IntegerVal(j)
	return i, nil
}

func (i IntegerVal) String() string {
	return fmt.Sprintf("%d", i)
}

type StringVal string

func (v StringVal) Set(s string) (Value, error) {
	v = StringVal(s)
	return v, nil
}

func (v StringVal) String() string {
	return string(v)
}
