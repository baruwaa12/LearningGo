package main

import (
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
		Bounds: pixel.R(0, 0, 100, 100),
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

	// Load the picture as a sprite
	sprite := pixel.NewSprite(pic, pic.Bounds())

	// Draw the sprite to the center of the window
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))


	for !win.Closed() {
		win.Update()
	}

}

func main () {
	pixelgl.Run(run)
}
