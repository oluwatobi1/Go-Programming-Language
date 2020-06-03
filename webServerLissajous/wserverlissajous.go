// An addition of the counter functionalities to the webserver.go
//NOTE
// Check wserver.go file for instruction on how to run this locally
// Assess counter functionality via "localhost:8080/count"

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// enter "localhost:8080/count" in any web browser on your system
//get the number of times requests has been made
func counter(write http.ResponseWriter, pth *http.Request) {
	mu.Lock()
	fmt.Fprintf(write, "count: %d\n", count)
	mu.Unlock()
}

var palette = []color.Color{color.Black, color.White}

const (
	firstColor  = 0
	secondColor = 1
)

func lissajous(out io.Writer) {
	const (
		cycle   = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), secondColor)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}
