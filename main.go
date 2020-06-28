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
	aspectRatio := 16.0 / 9.0
	imageWidth := 640
	imageHeight := int(float64(imageWidth) / aspectRatio)

	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := cartesian.NewPoint3(0, 0, 0)

	xDim := cartesian.NewVec3(viewportWidth, 0, 0)
	yDim := cartesian.NewVec3(0, viewportWidth, 0)
	halfXDim := cartesian.Div(&xDim, 2.0)
	halfYDim := cartesian.Div(&yDim, 2.0)
	focalLengthVec := cartesian.NewVec3(0, 0, focalLength)
	lowerLeftCorner := origin
	lowerLeftCorner = cartesian.Sub(&lowerLeftCorner, &halfXDim)
	lowerLeftCorner = cartesian.Sub(&lowerLeftCorner, &halfYDim)
	lowerLeftCorner = cartesian.Sub(&lowerLeftCorner, &focalLengthVec)

	fmt.Printf("lowerLeftCorner: %+v", lowerLeftCorner)

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
