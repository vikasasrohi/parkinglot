package model

type Gate struct {
	// ID     uint
	Number   uint
	Type     GateType
	Operator *Operator
}

type GateType string

const (
	Entry  GateType = "entry"
	Exit   GateType = "exit"
	Normal GateType = "normal"
)
