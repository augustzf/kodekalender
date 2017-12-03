package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// steganography, most likely?
func main() {
	file, err := os.Open("img.png")
	check(err)
	img, err := png.Decode(file)
	check(err)
	steg(img)
}

func steg(img image.Image) {
	var ch uint
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			b := uint(r & 0x01)
			var i = uint(x) % 8
			ch = ch | b<<i

			if i == 7 {
				// the solution is in the first complete bytes
				fmt.Print(string(ch))
				ch = 0
			}
		}

	}
}
