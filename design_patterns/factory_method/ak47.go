package main

type ak47 struct {
	gun
}

func newAK47() iGun {
	return &ak47{
		gun{
			"ak47",
			100,
		},
	}
}
