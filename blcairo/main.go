// Package main renders an image or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
)

//revive:disable:unused-parameter
const (
	// height = 270.0
	height = 540.0
	// height = 1080.0
	// height = 2160.0

	width       = height * render.Aspect16x9
	baseHeight  = 1080.0
	screenScale = height / baseHeight

	tau = blmath.Tau
	pi  = math.Pi
)

func main() {
	singleImage := true
	if singleImage {
		doImage()
	} else {
		doMovie()
	}
}

func doImage() {
	path := "out/image.png"
	render.Image(width, height, path, act1, 0)
	render.ViewImage(path)
}

func doMovie() {
	movie := render.NewMovie("movie", width, height, 30)
	movie.NewAct("act1", 240, act1, true, false)
	movie.CombineAll(true)
}

func act1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	context.SetLineWidth(0.5 * screenScale)
	for range 100 {
		context.MoveTo(random.FloatRange(0, width), random.FloatRange(0, height))
		context.LineTo(random.FloatRange(0, width), random.FloatRange(0, height))
	}
	context.Stroke()
}
