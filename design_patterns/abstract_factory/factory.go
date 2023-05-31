package main

func buildFactory(brand string) ifactory {
	switch brand {
	case "adidas":
		return &adidas{brand}
	case "nike":
		return &nike{brand}
	default:
		return &adidas{brand}
	}
}
