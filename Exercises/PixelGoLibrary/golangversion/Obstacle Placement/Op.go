package main

import (
	// "time"
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// PipePair structure which contains all its properties
type PipePair struct {
	FacingUp   	pixel.Picture
	FacingDown	pixel.Picture
	frameUp     pixel.Vec
	frameDown   pixel.Vec
	x,y         float64
}

func (pipes *PipePair) draw(win *pixelgl.Window) {
	// if dx = 100
	// dx := time.Since(pipeup.draw)
	
	// Loading the pipes as a sprite
	pipeup := pixel.NewSprite(pipes.FacingUp, pipes.FacingUp.Bounds())
	pipeup.Draw(win, pixel.IM.Moved(pipes.frameUp))

	pipedown := pixel.NewSprite(pipes.FacingDown, pipes.FacingDown.Bounds())
	pipedown.Draw(win, pixel.IM.Moved(pipes.frameDown))
}

// Function to load the picture
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {

	// Attributes for the test window.
	cfg := pixelgl.WindowConfig{
		Title:  "TestWindow",
		Bounds: pixel.R(0, 0, 1024, 582),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic1, err := loadPicture("City.png")
	if err != nil {
		panic(err)
	}

	pipeDownPic, err := loadPicture("pipe.png")
	if err != nil {
		panic(err)
	}
	
	pipeUpPic, err := loadPicture("pipedown.png")
	if err != nil {
		panic(err)
	}

	background  := pixel.NewSprite(pic1, pic1.Bounds())
	background.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	winRect := win.Bounds().Center()
	rectDown := pixel.Vec{winRect.X, win.Bounds().H()/4}
	rectUp := pixel.Vec{winRect.X, (win.Bounds().H()*0.75)}

	pipepair := PipePair{FacingUp: pipeUpPic, FacingDown: pipeDownPic, frameUp: rectUp, frameDown: rectDown}
	pipepair.draw(win)

	for !win.Closed() {
		win.Update()
	}
}



func main() {
	pixelgl.Run(run)
}