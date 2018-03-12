package fnnode

import (
	"errors"
)

type FnSpec interface {
	ID() string
}

// SpecFactory //////////////////////////////////////////////////

type SpecFactorySt struct {
	reg map[string]FnSpec
}

var SpecFactory = SpecFactorySt{
	reg: make(map[string]FnSpec),
}

func (factory SpecFactorySt) Register(spec FnSpec) error {
	//
	// TODO - erro out if 'id' already exists.
	//
	factory.reg[spec.ID()] = spec
	return nil
}

func (factory SpecFactorySt) lookup(id string) (FnSpec, error) {
	spec, ok := factory.reg[id]
	if !ok {
		return nil, errors.New("No FnSpec matching id.")
	}
	return spec, nil
}

// NilSpec - an empty Spec //////////////////////////////////////

type NilSpec struct{}

func (spec NilSpec) ID() string {
	return "nil-spec"
}

func MakeNilSpec() FnSpec {
	return NilSpec{}
}
