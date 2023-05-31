package main

type adidasShoes struct {
	logo string
	size int32
}

func (shoes *adidasShoes) setLogo(name string) {
	shoes.logo = name
}

func (shoes *adidasShoes) setSize(size int32) {
	shoes.size = size
}

func (shoes *adidasShoes) getLogo() string {
	return shoes.logo
}

func (shoes *adidasShoes) getSize() int32 {
	return shoes.size
}

type adidasShirt struct {
	logo string
	size int32
}

func (shoes *adidasShirt) setLogo(name string) {
	shoes.logo = name
}

func (shoes *adidasShirt) setSize(size int32) {
	shoes.size = size
}

func (shoes *adidasShirt) getLogo() string {
	return shoes.logo
}

func (shoes *adidasShirt) getSize() int32 {
	return shoes.size
}

type adidas struct {
	name string
}

func (a *adidas) makeShoes(size int32) ishoes {
	return &adidasShoes{
		logo: a.name,
		size: size,
	}
}

func (a *adidas) makeShirt(size int32) ishirt {
	return &adidasShirt{
		logo: a.name,
		size: size,
	}
}
