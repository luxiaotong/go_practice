package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	imgfile, err := os.Open("./sign.png")
	defer imgfile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	decodedImg, err := png.Decode(imgfile)
	img := image.NewRGBA(decodedImg.Bounds())

	cc := decodedImg.At(0, 0)
	//r, g, b, a := cc.RGBA()
	//fmt.Printf("color: %v, (%v, %v, %v, %v)\n", cc, r, g, b, a)
	size := img.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			c := decodedImg.At(x, y)
			if !similar(c, cc) {
				img.Set(x, y, decodedImg.At(x, y))
			}
		}
	}

	// I change some pixels here with img.Set(...)

	outFile, _ := os.Create("./output.png")
	defer outFile.Close()
	_ = png.Encode(outFile, img)
}

func similar(a, b color.Color) bool {
	r1, g1, b1, _ := a.RGBA()
	r2, g2, b2, _ := b.RGBA()
	if math.Abs((float64(r1)-float64(r2))/65535.0) < 0.05 && math.Abs((float64(g1)-float64(g2))/65535.0) < 0.05 &&
		math.Abs((float64(b1)-float64(b2))/65535.0) < 0.05 {
		return true
	}
	//fmt.Println(float64(r1-r2)/65535.0, "=", r2, r1-r2, r1, g1, b1)
	return false
}
