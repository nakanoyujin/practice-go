package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//クエリの取得
	cycles := r.URL.Query().Get("cycles")
	cyc, _ := strconv.Atoi(cycles)
	lissajous(w, cyc)
}

func lissajous(out io.Writer, cycle int) {
	var cycles = 5.0
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	if cycle != 0 {
		cycles = float64(cycle)
	}
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		//フレーム分だけ画像を生成
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, plt)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if rand.Int()%2 == 0 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 2)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}

var green = color.RGBA{0x00, 0xff, 0x00, 0xff}
var red = color.RGBA{0xff, 0x00, 0x00, 0xff}
var plt = color.Palette{color.Black, green, red}

const (
	//whiteIndex = 0
	greenIndex = 2
	blackIndex = 1
)
