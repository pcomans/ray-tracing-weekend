package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	imageWidth := 256
	imageHeight := 256

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	l := log.New(os.Stderr, "main: ", log.Lmicroseconds)

	for j := imageHeight - 1; j >= 0; j-- {
		l.Printf("Scanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / float64(imageWidth-1)
			g := float64(j) / float64(imageHeight-1)
			b := 0.25

			ir := uint8(255.999 * r)
			ig := uint8(255.999 * g)
			ib := uint8(255.999 * b)

			img.Set(i, j, color.RGBA{ir, ig, ib, 255})
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
