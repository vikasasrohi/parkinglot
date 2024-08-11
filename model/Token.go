package model

import "time"

type Token struct {
	ID        uint
	Number    uint
	Vehicle   *Vehicle
	Spot      *Spot
	Gate      *Gate
	Operator  *Operator
	EntryTime time.Time
}
