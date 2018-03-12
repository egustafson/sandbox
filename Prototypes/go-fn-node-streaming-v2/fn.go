package fnnode

import (
	"errors"
	_ "fmt"
)

type Msg interface {
	Body() interface{}
}

type FnProto func(spec FnSpec) (Fn, error)

type Fn func(in Msg) (out Msg)

func MakeUnityFn(spec FnSpec) (Fn, error) {
	return func(in Msg) (out Msg) {
		out = in
		return
	}, nil
}

type FnFactorySt struct {
	reg map[string]FnProto
}

var FnFactory = FnFactorySt{
	reg: make(map[string]FnProto),
}

func (fac FnFactorySt) Register(id string, proto FnProto) error {
	//
	// TODO - error out if 'id' already exists.
	//
	fac.reg[id] = proto
	return nil
}

func (fac FnFactorySt) Build(spec FnSpec) (Fn, error) {
	proto, ok := fac.reg[spec.ID()]
	if !ok {
		return nil, errors.New("No FnProto matching spec.")
	}
	return proto(spec)
}
