package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"

	"rsc.io/qr"
)

func main() {
	// fmt.Println("go-qrcode")
	// qr, _ := qrcode.New("http://bing.com", qrcode.Medium)
	// qr.BackgroundColor = color.Transparent
	// img := qr.Image(256)
	// b := img.Bounds()
	// m := image.NewRGBA(b)
	// draw.Draw(m, b, img, image.Point{0, 0}, draw.Over)
	// f, err := os.Create("image.png")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer f.Close()
	// if err := png.Encode(f, m); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println("go-qrcode")
	c, _ := qr.Encode("http://bing.com", qr.L)
	size := 256
	realSize := c.Size
	rect := image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{size, size}}
	bgColor := color.Transparent
	fgColor := color.Black
	p := color.Palette([]color.Color{bgColor, fgColor})
	img := image.NewPaletted(rect, p)
	fgClr := uint8(img.Palette.Index(fgColor))
	modulesPerPixel := float64(realSize) / float64(size)
	for y := 0; y < size; y++ {
		y2 := int(float64(y) * modulesPerPixel)
		for x := 0; x < size; x++ {
			x2 := int(float64(x) * modulesPerPixel)
			if c.Black(x2, y2) {
				pos := img.PixOffset(x, y)
				img.Pix[pos] = fgClr
			}
		}
	}
	var b bytes.Buffer
	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	_ = encoder.Encode(&b, img)
	_ = ioutil.WriteFile("rsc.png", b.Bytes(), os.FileMode(0644))
}
