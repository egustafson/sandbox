package tunable

import "fmt"

// This package is heavily inspired by the golang flag and spf13/pflags packages.

type TunableSet struct {
	name     string
	tunables map[string]*Tunable
	defaults map[string]any
}

type Tunable struct {
	ID         string
	Overriden  bool
	Value      Value
	DefaultVal Value
	InitialVal Value
}

func NewTunableSet(name string) *TunableSet {
	return &TunableSet{
		name:     name,
		tunables: make(map[string]*Tunable),
	}
}

func (ts *TunableSet) Exists(name string) bool {
	_, ok := ts.tunables[name]
	return ok
}

func (ts *TunableSet) SetBool(name string, value bool) error {
	if t, exists := ts.tunables[name]; exists {
		return t.Value.SetValue(value)
	}
	return notExistError(name)
}

func (ts *TunableSet) DefineBoolVar(p *bool, name string, value bool) {
	ts.DefineVar(newBoolValue(value, p), name)
}

func (ts *TunableSet) SetInt(name string, value int) error {
	if t, exists := ts.tunables[name]; exists {
		return t.DefaultVal.SetValue(value)
	}
	return notExistError(name)
}

func (ts *TunableSet) DefineIntVar(p *int, name string, value int) {
	ts.DefineVar(newIntValue(value, p), name)
}

func (ts *TunableSet) DefineVar(value Value, id string) {
	tunable := &Tunable{
		ID:         id,
		Overriden:  false,
		Value:      value,
		DefaultVal: value.Clone(),
		InitialVal: value.Clone(),
	}
	_, alreadythere := ts.tunables[id]
	if alreadythere {
		fatalError(fmt.Errorf("tunable already defined: %s", id))
		return
	}
	ts.tunables[id] = tunable
	defaultVal, defaultExists := ts.defaults[id]
	if defaultExists {
		ts.SetDefault(id, defaultVal)
	}
}

func (ts *TunableSet) SetDefault(name string, defaultVal any) error {
	if t, exists := ts.tunables[name]; exists {
		if err := t.DefaultVal.SetValue(defaultVal); err != nil {
			return err
		}
		if !t.Overriden {
			t.Value.SetValue(defaultVal)
		}
		return nil
	}
	return notExistError(name)
}

func (ts *TunableSet) Set(name string, val any) error {
	if t, exists := ts.tunables[name]; exists {
		if err := t.Value.SetValue(val); err != nil {
			return err
		}
		t.Overriden = true
		return nil
	}
	return notExistError(name)
}
