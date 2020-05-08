package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var palette = map[complex128]color.Color{
	complex(1, 0):  color.RGBA{R: 255},
	complex(-1, 0): color.RGBA{G: 255},
	complex(0, 1):  color.RGBA{B: 255},
	complex(0, -1): color.RGBA{},
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -0.25, -0.25, +0.25, +0.25
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	const iterations = 255
	const contrast = 15
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		for root, rootColor := range palette {
			if cmplx.Abs(z-root) < 1e-6 {
				r, g, b, _ := rootColor.RGBA()
				return color.RGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: 255 - contrast*i,
				}
			}
		}
	}

	return color.Black
}
