package model

type Floor struct {
	ID              uint
	Name            string
	Number          uint
	AllowedVehicles []VehicleType
	Spots           []Spot
}
