// Creating a game window

package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)


func run() {
		cfg := pixelgl.WindowConfig{

			// Defines Window properties
			Title: "Pixel rocks",
			Bounds: pixel.R(0, 0, 1024, 728),
			VSync: true,
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		win.Clear(colornames.Skyblue)
		for !win.Closed() {
				win.Update()
		}
}

func main() {
		pixelgl.Run(run)
}