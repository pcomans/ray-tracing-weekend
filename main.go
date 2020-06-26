package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"phpp.me/go_practice/cartesian"
)

func main() {
	imageWidth := 256
	imageHeight := 256

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	l := log.New(os.Stderr, "main: ", log.Lmicroseconds)

	for j := imageHeight - 1; j >= 0; j-- {
		l.Printf("Scanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			c := cartesian.NewColor(
				255.999*(float64(i)/float64(imageWidth-1)),
				255.999*(float64(j)/float64(imageHeight-1)),
				255.999*0.25,
			)

			d := cartesian.NewColor(128.0, 128.0, 128.0)
			fmt.Printf("%v", d)

			img.Set(i, j, color.RGBA{uint8(c.X), uint8(c.Y), uint8(c.Z), 255})
		}
	}

	f, err := os.Create("out.png")
	if err != nil {
		fmt.Printf("Error: %d\n", err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		fmt.Printf("Error: %d\n", err)
	}
}
