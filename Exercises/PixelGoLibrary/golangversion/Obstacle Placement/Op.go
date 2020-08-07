package main

import (
	"time"
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
	x			float64
	yUp 		float64
	yDown 		float64
	lastDrawn	time.Time
}

func (pipes *PipePair) draw(win *pixelgl.Window) {

	dt := time.Since(pipes.lastDrawn).Milliseconds();
	
	if dt >= 10 {
		pipes.x = pipes.x - 1 
		pipes.lastDrawn = time.Now()
	}

	newFrameUpVector := pixel.Vec{X: pipes.x, Y: pipes.yUp}
	newFrameDownVector := pixel.Vec{X: pipes.x, Y: pipes.yDown}
	
	pipeup := pixel.NewSprite(pipes.FacingUp, pipes.FacingUp.Bounds())
	pipeup.Draw(win, pixel.IM.Moved(newFrameUpVector))

	pipedown := pixel.NewSprite(pipes.FacingDown, pipes.FacingDown.Bounds())
	pipedown.Draw(win, pixel.IM.Moved(newFrameDownVector))
	
	
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

	winRect := win.Bounds().Center()
	rectDown := pixel.Vec{winRect.X, win.Bounds().H()/4}
	rectUp := pixel.Vec{winRect.X, (win.Bounds().H()*0.75)}

	pipepair := PipePair{FacingUp: pipeUpPic, FacingDown: pipeDownPic, x: rectUp.X, yUp: rectUp.Y, yDown: rectDown.Y,  lastDrawn: time.Now()}

	for !win.Closed() {
		background.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		pipepair.draw(win)
		win.Update()
	}
}



func main() {
	pixelgl.Run(run)
}