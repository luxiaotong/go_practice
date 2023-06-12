package main

type iglooBuilder struct {
	windowType string
	doorType   string
	floorNum   int32
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (n *iglooBuilder) setDoorType() {
	n.doorType = "Snow Door"
}

func (n *iglooBuilder) setWindowType() {
	n.windowType = "Snow Window"
}

func (n *iglooBuilder) setFloorNum() {
	n.floorNum = 1
}

func (n *iglooBuilder) getHouse() *House {
	return &House{
		windowType: n.windowType,
		doorType:   n.doorType,
		floorNum:   n.floorNum,
	}
}
