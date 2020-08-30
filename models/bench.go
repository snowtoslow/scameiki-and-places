package models

type Bench struct {
	ID int
	Geolocation string
	Photo []byte
}

type Benches []Bench
