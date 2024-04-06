// Package main renders an image, gif or video
package main

import (
	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
)

func main() {

	renderTarget := target.Video

	switch renderTarget {
	case target.Image:
		render.Image(800, 800, "out/out.png", scene1, 0.0)
		render.ViewImage("out/out.png")
		break

	case target.Video:
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 60)
		program.RenderVideo("out/frames", "out/out.mp4")
		render.PlayVideo("out/out.mp4")
		break
	}
}

//revive:disable-next-line:unused-parameter
func scene1(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.Save()
	context.TranslateCenter()
	context.DrawAxes(0.25)
	r := blmath.LoopSin(percent, 50, width/2)
	sphere := cairo.NewSphere(0, 0, r, 1, 0, 0)
	sphere.SetShadowColor(0.2, 0, 0)
	sphere.Draw(context)
	context.Restore()
}
