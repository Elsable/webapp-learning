package main

import (
	"net/http"
	"log"
	"io"
	"math/rand"
	"image/gif"
	"image"
	"math"
	"image/color"
)

var palette2 = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0x00},
}

const (
	whiteIndex_ = 0
	blackIndex_ = 1
	greenIndex_ = 2
)

func main()  {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous4(w)
}

func lissajous4(out io.Writer)  {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex_)
		}
		phase += 1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
