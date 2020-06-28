package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	cc "phpp.me/go_practice/cartesian"
)

func rayColor(r cc.Ray) cc.Color {
	unitDirection := cc.UnitVector(r.Direction)
	t := 0.5*unitDirection.Y + 1.0
	return cc.Add(cc.Mul(cc.NewColor(1, 1, 1), 1.0-t), cc.Mul(cc.NewColor(0.5, 0.7, 1), t))
}

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 640
	imageHeight := int(float64(imageWidth) / aspectRatio)

	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := cc.NewPoint3(0, 0, 0)

	xDim := cc.NewVec3(viewportWidth, 0, 0)
	yDim := cc.NewVec3(0, viewportWidth, 0)

	lowerLeftCorner := cc.Sub(origin, cc.Div(xDim, 2.0), cc.Div(yDim, 2.0), cc.NewVec3(0, 0, focalLength))

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	l := log.New(os.Stderr, "main: ", log.Lmicroseconds)

	for j := imageHeight - 1; j >= 0; j-- {
		l.Printf("Scanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)

			r := cc.NewRay(origin, cc.Sub(cc.Add(lowerLeftCorner, cc.Mul(xDim, u), cc.Mul(yDim, v)), origin))
			c := rayColor(r)

			img.Set(i, j, color.RGBA{uint8(255.999 * c.X), uint8(255.999 * c.Y), uint8(255.999 * c.Z), 255})
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
