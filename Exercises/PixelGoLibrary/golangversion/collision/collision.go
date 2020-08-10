package main 

import (
		"fmt"
		"image"
		"os"
		_ "image/png"
		"github.com/faiface/pixel/pixelgl"
		"github.com/faiface/pixel"
)



// Load the birds picture
func loadPictureBird(path string) (pixel.Picture, error) {
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

	// Window properties
	cfg := pixelgl.WindowConfig{
		Title: "TestWindow",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPictureBird("wingup.png")
	if err != nil {
		panic(err)
	}

	pipepic, err := loadPictureBird("pipe.png")
	if err != nil {
		panic(err)
	}

	// Load the picture as a sprite
	bird := pixel.NewSprite(pic, pic.Bounds())
	pipe := pixel.NewSprite(pipepic, pipepic.Bounds())

	windowBound := win.Bounds().Center()
	pipeVector := pixel.Vec{X: windowBound.X + 100, Y: windowBound.Y} 
	// // Draw the sprite to the center of the window
	bird.Draw(win, pixel.IM.Moved(windowBound))
	pipe.Draw(win, pixel.IM.Moved(pipeVector))
	

	fmt.Println("Pipe Rect %T", pipeVector)
	fmt.Println("Bird Rect %T", windowBound)

	pipeX := pipeVector.X - 80/2
	pipeY1 := pipeVector.Y + (150/2)
	pipeY :=  pipeVector.Y - (150/2)

	for !win.Closed() {
		win.Update()
	}
}

func main( ) {
	pixelgl.Run(run)
}