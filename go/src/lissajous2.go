package main

import (
	"math/rand"
	"time"
	"io"
	"os"
	"image/gif"
	"image"
	"image/color"
	"math"
)

func main()  {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous2(os.Stdout)
}

func lissajous2(out io.Writer)  {
	var palette2 = []color.Color{
		color.Black,
		color.RGBA{0x00, 0xff, 0x00, 0xff},
	}
	const (
		greenIndex = 1
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3
	animation := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+5), size+int(y*size+5), greenIndex)
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)
	}
	gif.EncodeAll(out, &animation)
}