package main

type normalBuilder struct {
	windowType string
	doorType   string
	floorNum   int32
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (n *normalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *normalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
}

func (n *normalBuilder) setFloorNum() {
	n.floorNum = 2
}

func (n *normalBuilder) getHouse() *House {
	return &House{
		windowType: n.windowType,
		doorType:   n.doorType,
		floorNum:   n.floorNum,
	}
}
