package main

import (
	"image/color"
	"net/http"
	"log"
	"io"
	"image/gif"
	"math/rand"
	"image"
	"math"
	"strconv"
)

type LissajousParams struct {
	Cycles int
	Resolution float64
	Size int
	NFrames int
	Delay int
}


var lissajousPalette = [] color.Color{
	color.White,
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
}

func main()  {
	http.HandleFunc("/", handlerLissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handlerLissajous(w http.ResponseWriter, r *http.Request) {
	var params LissajousParams
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	params.Cycles = 5
	params.Resolution = 0.01
	params.Size = 100
	params.NFrames = 64
	params.Delay = 8
	for key, value := range r.Form {
		if key == "cycles" {
			params.Cycles, _ = strconv.Atoi(value[0])
		} else if key == "res" {
			params.Resolution, _ = strconv.ParseFloat(value[0], 32)
		} else if key == "size" {
			params.Size, _ = strconv.Atoi(value[0])
		} else if key == "nframes" {
			params.NFrames, _ = strconv.Atoi(value[0])
		} else if key == "delay" {
			params.Delay, _ = strconv.Atoi(value[0])
		}
	}
	serverLissajous(w, params)
}

func serverLissajous(out io.Writer, params LissajousParams) {
	const (
		blackIndex = 1
	)
	animation := gif.GIF{LoopCount: params.NFrames}
	freq := rand.Float64() * 3.0
	phase := 0.0
	for i := 0; i < params.NFrames; i++ {
		rect := image.Rect(0, 0, 2*params.Size+1, 2*params.Size+1)
		img := image.NewPaletted(rect, lissajousPalette)
		for t := 0.0; t < float64(params.Cycles)* 2 *math.Pi; t+= params.Resolution {
			x := math.Sin(t)
			y := math.Sin(t*freq+phase)
			img.SetColorIndex(params.Size+int(x*float64(params.Size)+5), params.Size+int(y*float64(params.Size)+5),
				              blackIndex)
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, params.Delay)
		animation.Image = append(animation.Image, img)
	}
	gif.EncodeAll(out, &animation)
}