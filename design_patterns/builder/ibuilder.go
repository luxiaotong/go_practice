package main

type ibuilder interface {
	setDoorType()
	setWindowType()
	setFloorNum()
	getHouse() *House
}

type House struct {
	windowType string
	doorType   string
	floorNum   int32
}
