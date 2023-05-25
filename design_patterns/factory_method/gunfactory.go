package main

func getGun(typ string) iGun {
	switch typ {
	case "ak47":
		return newAK47()
	case "musket":
		return newmusket()
	default:
		return newAK47()
	}
}
