package model

type Spot struct {
	ID          uint
	Number      uint
	VehicleType VehicleType
	Floor       *Floor
	Status      SpotStatus
}

type SpotStatus string

const (
	Empty    SpotStatus = "empty"
	Booked   SpotStatus = "booked"
	Occupied SpotStatus = "occupied"
)
