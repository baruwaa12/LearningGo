package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

)

// Main
func main() {
	//	defer profile.Start(profile.CPUProfile).Stop()
	pixelgl.Run(run)
}

func run() {

	global.gVariableConfig.LoadConfiguration()

	cfg := pixelgl.WindowConfig{
		Title:       "PacMan",
		Bounds:      pixel.R(0, 0, float64(800), float64(600)),
		VSync:       false,
		Undecorated: false,
	}
	
	gWin, err := pixelgl.NewWindow(cfg)

	Debug("W:", gWin.Bounds())

	gWin.SetCursorVisible(false)

	if err != nil {
		panic(err)
	}

	centerWindow(gWin)
	global.gWin = gWin

	setup()

	gameLoop()

}

func setup() {
	global.gFont.create()
	var selected int32

	selected = 2
	// Create Menu
	global.gCanvas =   pixelgl.NewCanvas(pixel.R(0, 0, 1, 1))
	global.gCanvas.SetUniform("uSelected", selected)
	global.gCanvas.SetUniform("uTime", &global.uTime)
	global.gCanvas.SetFragmentShader(fragmentShaderMenuItem)

	global.gWin.SetSmooth(false)

	global.gWin.Canvas().SetFragmentShader(fragmentShaderFullScreen)
}

func gameLoop() {

		last := time.Now()
		frameDt := 0.0
	
		
		for !global.gWin.Closed() {
			dt := time.Since(last).Seconds()
			frameDt += dt
			last = time.Now()

			global.gCanvas.Clear(pixel.RGBA{R: 0, G: 0, B: 0, A: 0})
			global.gFont.writeToCanvas("Hello", global.gCanvas)
			offsetX := (float64(global.gVariableConfig.WindowWidth) / 3) / 2 //- m.logo.canvas.Bounds().Max.X/2
			offsetY := (float64(global.gVariableConfig.WindowHeight) / 3) - global.gCanvas.Bounds().Max.Y
			logoScale := float64(global.gVariableConfig.WindowWidth) / float64(global.gVariableConfig.WindowHeight)
			global.gCanvas.Draw(global.gWin, pixel.IM.Scaled(pixel.ZV, logoScale).Moved(pixel.V(0+offsetX, 0+offsetY)))
			
			global.gWin.Update()
		}
}