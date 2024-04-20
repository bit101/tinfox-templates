// Package main renders an image, gif or video
package main

import (
	"github.com/bit101/bitlib/bllog"
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
)

func main() {
	bllog.InitProjectLog("project.log")
	defer bllog.CloseLog()

	renderTarget := target.Video

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/out.png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 60)
		program.RenderAndPlayVideo("out/frames", "out/out.mp4")
	}
}

var (
	points geom.PointList
)

func init() {
	points = geom.RandomPointList(1000, 0, 0, 400, 400)
}

//revive:disable-next-line:unused-parameter
func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	p := points.Clone()
	r := blmath.LoopSin(percent, 0, 100)
	p.Randomize(r, r)
	context.Points(p, 1)
}
