package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	canvas := image.NewRGBA(image.Rect(0, 0, 2*width, 2*height))
	for py := 0; py < 2*height; py++ {
		y := ymin + (ymax-ymin)*float64(py)/(2*height)
		for px := 0; px < 2*width; px++ {
			x := xmin + (xmax-xmin)*float64(px)/(2*width)
			z := complex(x, y)
			canvas.Set(px, py, mandelbrot(z))
		}
	}

	sampled := supersample(canvas)

	if err := png.Encode(os.Stdout, sampled); err != nil {
		log.Fatal(err)
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				R: 255 - 15*n,
				G: 255 - 25*n,
				B: 255 - 35*n,
				A: 255,
			}
		}
	}
	return color.Black
}

func averageColor(samples []color.Color) color.Color {
	var rr, gg, bb, aa float64
	for _, sample := range samples {
		r, g, b, a := sample.RGBA()
		rr += float64(r)
		gg += float64(g)
		bb += float64(b)
		aa += math.Sqrt(float64(a))
	}
	count := float64(len(samples))
	return color.RGBA{
		R: uint8(rr / aa),
		G: uint8(gg / aa),
		B: uint8(bb / aa),
		A: uint8(aa / count),
	}
}

func supersample(source *image.RGBA) *image.RGBA {
	size := source.Rect.Size()
	width, height := size.X/2, size.Y/2
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			sampled := averageColor([]color.Color{
				source.At(2*x, 2*y),
				source.At(2*x+1, 2*y),
				source.At(2*x, 2*y+1),
				source.At(2*x+1, 2*y+1),
			})
			img.Set(x, y, sampled)
		}
	}

	return img
}
