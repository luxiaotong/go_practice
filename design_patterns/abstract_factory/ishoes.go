package main

type ishoes interface {
	setSize(size int32)
	setLogo(name string)
	getSize() int32
	getLogo() string
}
