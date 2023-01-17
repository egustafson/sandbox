package tunable

import "fmt"

// This package is heavily inspired by the golang flag and spf13/pflags packages.

type TunableSet struct {
	name     string
	tunables map[string]*Tunable
}

type Tunable struct {
	ID        string
	Overriden bool
	Value     Value
	Default   Value
}

func NewTunableSet(name string) *TunableSet {
	return &TunableSet{
		name:     name,
		tunables: make(map[string]*Tunable),
	}
}

func (t *TunableSet) SetBool(name string, value bool) error {
	//
	// TODO: implement this and the general Set method
	//
	return nil
}

func (t *TunableSet) DefineBoolVar(p *bool, name string, value bool) {
	t.DefineVar(newBoolValue(value, p), name)
}

func (t *TunableSet) SetInt(name string, value int) error {
	//
	// TODO: implement this and the general Set method
	//
	return nil
}

func (t *TunableSet) DefineIntVar(p *int, name string, value int) {
	t.DefineVar(newIntValue(value, p), name)
}

func (t *TunableSet) DefineVar(value Value, id string) {
	tunable := &Tunable{
		ID:        id,
		Overriden: false,
		Value:     value,
		Default:   value,
	}
	_, alreadythere := t.tunables[id]
	if alreadythere {
		// IMPROVE - there's room to improve handling this corner case
		panic(fmt.Errorf("tunable already defined: %s", id))
	}
	t.tunables[id] = tunable
}
