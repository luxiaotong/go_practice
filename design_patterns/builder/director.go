package main

type director struct {
	builder ibuilder
}

func newDirector(b ibuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) buildHouse() *House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setFloorNum()
	return d.builder.getHouse()
}
