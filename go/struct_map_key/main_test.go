package main_test

import (
	"fmt"
	"testing"
)

type DataStruct struct {
	Id   int
	Name string
}

type OtherStruct struct {
	Id   int
	Name string
}

type DataStructMap map[any]*DataStruct

func TestStructKey(t *testing.T) {

	var mapOfData DataStructMap = make(DataStructMap)

	setData(mapOfData, 1)
	setData(mapOfData, 2)
	setData(mapOfData, 3)

	var k2 any = OtherStruct{}
	mapOfData[k2] = &DataStruct{Id: 99, Name: "error case"} // <-- attempt to create an error
	fmt.Printf("k2: %v\n", k2)

	magic := getData(mapOfData)

	if magic.Id != 3 {
		t.Error("returned struct does not match original")
	}
}

func getData(m DataStructMap) *DataStruct {
	return m[DataStruct{}]
}

func setData(m DataStructMap, id int) {
	var k any = DataStruct{}
	m[k] = &DataStruct{Id: id, Name: "anyname"}
	fmt.Printf("k:  %v\n", k)
}
