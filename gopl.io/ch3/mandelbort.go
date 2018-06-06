package ch3

import (
	"image/color"
	"math/cmplx"
	"image"
	"image/png"
	"io"
	"sync"
)

func Mm(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		// 会有迭代变量的问题
		y := float64(py)/height*(ymax-ymin) + ymin
		wg.Add(1)

		go func(gy float64, gpy int) {
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, gy)

				img.Set(px, gpy, Mandelbrot(z))
			}
			wg.Done()
		}(y, py)

		//for px := 0; px < width; px++ {
		//	x := float64(px)/width*(xmax-xmin) + xmin
		//	z := complex(x, y)
		//
		//	img.Set(px, py, Mandelbrot(z))
		//}

	}

	wg.Wait()
	png.Encode(w, img)
}

func Mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}

	return color.Black
}
