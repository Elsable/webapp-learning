package main

import (
	"math/rand"
	"time"
	"os"
	"io"
	"image/color"
	"image/gif"
	"image"
	"math"
)

func main()  {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous3(os.Stdout)
}

func lissajous3(out io.Writer)  {
	var palette3 = []color.Color{
		color.Black,
		color.White,
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x88, 0x88, 0x00, 0xff},
		color.RGBA{0x00, 0x88, 0x88, 0xff},
		color.RGBA{0x88, 0x00, 0x88, 0xff},
	}
	const (
		cycles = 8
		res = 0.001
		size = 100
		nframes = 128
		delay = 8
	)
	freq := rand.Float64() * 3.0
	animation := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette3)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			idx := rand.Int() % 7 + 1
			img.SetColorIndex(size+int(size*x+8), size+int(size*y+8), uint8(idx))
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)
	}
	gif.EncodeAll(out, &animation)
}