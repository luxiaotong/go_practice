package main

type iGun interface {
	setName(name string)
	setPower(power int32)
	getName() string
	getPower() int32
}

type gun struct {
	name  string
	power int32
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int32) {
	g.power = power
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int32 {
	return g.power
}
