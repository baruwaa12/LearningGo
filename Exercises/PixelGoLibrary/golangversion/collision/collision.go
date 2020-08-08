package main

import (
	"time"
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)





// CB = collision box
// PipePair structure which contains all its properties
type PipePair struct {
	FacingUp   	 pixel.Picture
	FacingDown	 pixel.Picture	
	x			 float64
	yUp 		 float64
	yDown 		 float64
	lastDrawn	 time.Time
	FacingUpCB   pixel.Rect
	FacingDownCB pixel.Rect
	FrameWidth   float64
}

func (pipes *PipePair) draw(win *pixelgl.Window) {


	// Takes the time since the last drawn pipe on screen
	dt := time.Since(pipes.lastDrawn).Milliseconds();
	
	// If the duration is over 10 milliseconds 
	// Pipe moves across the screen 1 pixel.
	if dt >= 10 {
		pipes.x = pipes.x - 1 
		pipes.lastDrawn = time.Now()
	}

	// If the pipe is at 20 pixels away from x = 0 
	// The pipe moves back to the x = 1300 of the screen
	if pipes.x < 20 {
		pipes.x = 1300
	}

	// Defines the position of where to place the pipes on screen
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

	// Create frames around the images as a hitbox
	var frames []pixel.Rect
	for x := 0.0; x+FrameWidth <= img.Bounds().Max.X; x += FrameWidth {
		frames = append(frames, pixel.R(
			x,
			0,
			x + FrameWidth,
			img.Bounds().H(),
		))
	}
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

	x := win.Bounds().W()

	// New instances of the pipepair using PipePairs struct
	pipepair := PipePair{FacingUp: pipeUpPic, FacingDown: pipeDownPic, x: x/2, yUp: rectUp.Y, yDown: rectDown.Y,  lastDrawn: time.Now()}
	pipepair2 := PipePair{FacingUp: pipeUpPic, FacingDown: pipeDownPic, x: x, yUp: rectUp.Y, yDown: rectDown.Y,  lastDrawn: time.Now()}


	// While the window is not closed, after every refresh the pipepairs should be drawn at its next position.
	for !win.Closed() {
		background.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		pipepair.draw(win)
		pipepair2.draw(win)
		win.Update()
	}
}



func main() {
	pixelgl.Run(run)
}