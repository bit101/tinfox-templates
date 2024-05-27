// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "${PROJECT_DIR}"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 180)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau = blmath.Tau
	pi  = math.Pi
)

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.Sphere(400, 10, 20, true, true)

	shape.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	shape.Stroke()

	// context.GaussianBlur(20)
	// context.SetLineWidth(1)
	// shape.Stroke()
}
