package main

import (
	"fmt"
	"time"
)

type Timestamp time.Time

func Now() Timestamp {
	return Timestamp(time.Now())
}

func NewTimeStamp(t time.Time) Timestamp {
	return Timestamp(t)
}

func (t Timestamp) Time() time.Time {
	return time.Time(t)
}

func (t Timestamp) String() string {
	return time.Time(t).UTC().Format(time.RFC3339)
}

func (t Timestamp) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

type Demo struct {
	ID int       `json:"id"`
	At Timestamp `json:"at"`
}

func (d *Demo) String() string {
	return fmt.Sprintf("{ id: %d, at: %s }", d.ID, d.At.String())
}

func main() {
	// do nothing
}
