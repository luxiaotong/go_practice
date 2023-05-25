package main

type musket struct {
	gun
}

func newmusket() iGun {
	return &musket{
		gun{
			"musket",
			100,
		},
	}
}
