package main

type nikeShoes struct {
	logo string
	size int32
}

func (shoes *nikeShoes) setLogo(name string) {
	shoes.logo = name
}

func (shoes *nikeShoes) setSize(size int32) {
	shoes.size = size
}

func (shoes *nikeShoes) getLogo() string {
	return shoes.logo
}

func (shoes *nikeShoes) getSize() int32 {
	return shoes.size
}

type nikeShirt struct {
	logo string
	size int32
}

func (shoes *nikeShirt) setLogo(name string) {
	shoes.logo = name
}

func (shoes *nikeShirt) setSize(size int32) {
	shoes.size = size
}

func (shoes *nikeShirt) getLogo() string {
	return shoes.logo
}

func (shoes *nikeShirt) getSize() int32 {
	return shoes.size
}

type nike struct {
	name string
}

func (n *nike) makeShoes(size int32) ishoes {
	return &nikeShoes{
		logo: n.name,
		size: size,
	}
}

func (n *nike) makeShirt(size int32) ishirt {
	return &nikeShirt{
		logo: n.name,
		size: size,
	}
}
