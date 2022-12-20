package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/google/uuid"
)

type Object struct {
	TS time.Time `yaml:"ts"`
	ID uuid.UUID `yaml:"id"`
}

func NewObject() Object {
	<-time.After(time.Duration((10 + rand.Intn(100))) * time.Microsecond)
	return Object{
		TS: time.Now().UTC(),
		ID: uuid.New(),
	}
}

type ByTime []Object

// static check: struct ByTime implements sort.Interface
var _ = sort.Interface(ByTime{})

func (objs ByTime) Len() int           { return len(objs) }
func (objs ByTime) Less(i, j int) bool { return objs[i].TS.Before(objs[j].TS) }
func (objs ByTime) Swap(i, j int)      { objs[i], objs[j] = objs[j], objs[i] }

func main() {

	size := 5
	objList := make([]Object, size)

	for ii := 0; ii < size; ii++ {
		objList[(size-1)-ii] = NewObject() // <- insert in reverse
	}

	fmt.Println("----- reversed:")
	for _, obj := range objList {
		fmt.Printf("%v\n", obj)
	}

	sort.Sort(ByTime(objList)) // <- sort the list of objects

	fmt.Println("----- sorted:")
	for _, obj := range objList {
		fmt.Printf("%v\n", obj)
	}

	fmt.Println("done.")
}
