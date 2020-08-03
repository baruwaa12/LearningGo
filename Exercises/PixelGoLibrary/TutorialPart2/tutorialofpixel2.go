package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig {
			Title: "Pixel Rocks!",
			Bounds: pixel.R(0, 0, 1024, 768),
			VSync: true,

	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}


	// imd represents a newly drawn image
	imd := imdraw.New(nil)

	// color
	imd.Color = colornames.Blueviolet

	imd.EndShape = imdraw.RoundEndShape

	// Push takes vectors representing the position points.
	imd.Push(pixel.V(100, 100), pixel.V(700, 100))
	imd.Push(pixel.V(100, 500), pixel.V(700, 500))

	// Line draws a line between all vectors.
	imd.Line(30)


	for !win.Closed() {
			win.Clear(colornames.Aliceblue)
			imd.Draw(win)
			win.Update()
	}

}

func main() {
	pixelgl.Run(run)
}