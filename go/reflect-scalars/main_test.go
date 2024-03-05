package main

import (
	"reflect"
	"testing"
	"time"
)

type uberString string

func TestStringVsInt(t *testing.T) {
	i := 42
	var s = "string"
	var us interface{} = uberString("uberString")

	if reflect.TypeOf(s).Kind() != reflect.String {
		t.Errorf("s is not a string")
	}

	if reflect.TypeOf(s).Kind() != reflect.String {
		t.Error("us is not a string")
	}

	if reflect.TypeOf(i).Kind() == reflect.String {
		t.Error("i is a string ?!?!")
	}

	_ = i
	_ = s
	_ = us
}

func ReturnAs(v interface{}) string {
	x, ok := v.(string)
	if ok {
		return x
	}
	return ""
}


func TestStringVsIntAlternative(t *testing.T) {

	if len(ReturnAs("abc")) != 3 {
		t.Error("len of 'abc' is not 3")
	}

	if len(ReturnAs(123)) != 0 {
		t.Error("len of integer is not 0")
	}

}

func IsInteger(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

func IsInt(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Int
}

func IsDuration(v interface{}) bool {
	const d time.Duration = 0
	return reflect.TypeOf(v).Kind() == reflect.TypeOf(d).Kind()
}

func TestIntVsDuration(t *testing.T) {
	if IsInt(0) != true {
		t.Error("0 is not an Integer ?!?!")
	}
	if IsDuration(1) == true {
		t.Error("1 is a Duration ?!?!?")
	}
	if IsInt(time.Minute) == true {
		t.Error("time.Minute is an Int ?!?!?")
	}
	if IsInteger(int32(1)) != true {
		t.Error("int32 != Integer - we've got a problem")
	}
	if IsInteger(time.Hour) == true {
		t.Error("time.Hour is an Integer ?!?!?!")
	}
}
