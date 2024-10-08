package model

type Vehicle struct {
	ID    uint
	Plate string
	Type  VehicleType
}

type VehicleType string

const (
	CarVehicle  VehicleType = "car"
	BikeVehicle VehicleType = "bike"
)
