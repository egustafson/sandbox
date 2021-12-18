package main

type HealthState int

const (
	Ok HealthState = iota
	Unhealthy
	Unknown
)

var healthStateStrings = map[HealthState]string{
	Ok:        "ok",
	Unhealthy: "unhealthy",
	Unknown:   "unknown",
}

func (s HealthState) String() string {
	return healthStateStrings[s]
}

func (s HealthState) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}
