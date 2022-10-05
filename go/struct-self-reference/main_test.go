package main_test

import (
	"fmt"
	"reflect"
	"testing"
)

type RecordA struct {
	ID   int
	Name string
}

type RecordB struct {
	ID    int
	Value int
}

type RecordC struct {
	ID    int
	Value float32
	Note  string
}

type RecordD struct {
	RecordA // <-- type extension
	Value   float32
}

type RecordE = RecordA // <-- type alias

// --  map of data  ------------------------------

type DataMap map[any]any

func TestTypeReference(t *testing.T) {

	var mapOfData = make(DataMap)
	mapOfData.updateA(RecordA{ID: 1, Name: "record-a"})
	mapOfData.updateB(RecordB{ID: 2, Value: 999})

	fmt.Printf("len(mapOfData) = %d, expect: 2\n", len(mapOfData))

	mapOfData.updateC(RecordC{ID: 3, Value: 3.14, Note: "record-c"})
	mapOfData.updateD(RecordD{RecordA{ID: 4, Name: "record-d"}, 1.41})
	mapOfData.updateE(RecordE{ID: 9, Name: "e-aliased-record-a"}) // <- aliased type = orig type

	fmt.Printf("len(mapOfData) = %d, expect: 5\n", len(mapOfData))

	for ii, k := range mapOfData.keys() {
		fmt.Printf("%d: %v - (%s)\n", ii, k, reflect.TypeOf(k).Name())
	}

	var a RecordA = mapOfData.getA()
	fmt.Printf("a: %v\n", a)

	mapOfData.updateA(RecordA{ID: 10, Name: "record-a-num-2"})
	mapOfData.updateC(RecordC{ID: 13, Value: 3.14, Note: "record-c-num2"})
	mapOfData.updateD(RecordD{RecordA{ID: 14, Name: "record-d-num2"}, 1.41})

	fmt.Printf("len(mapOfData) = %d, expect: 4\n", len(mapOfData))

	for ii, k := range mapOfData.keys() {
		fmt.Printf("%d: %v - (%s)\n", ii, k, reflect.TypeOf(k).Name())
	}

	a = mapOfData.getA()
	fmt.Printf("a: %v\n", a)
}

func (m DataMap) updateA(d RecordA) { m[RecordA{}] = d }
func (m DataMap) updateB(d RecordB) { m[RecordB{}] = d }
func (m DataMap) updateC(d RecordC) { m[RecordC{}] = d }
func (m DataMap) updateD(d RecordD) { m[RecordD{}] = d }
func (m DataMap) updateE(d RecordE) { m[RecordE{}] = d }

func (m DataMap) getA() RecordA {
	if v, ok := m[RecordA{}].(RecordA); ok {
		return v
	}
	return RecordA{}
}

func (m DataMap) keys() []any {
	keys := make([]any, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}
