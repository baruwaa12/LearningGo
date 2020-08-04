// Scrolling backrgound test

package main

import (
	"golang.org/x/image/colornames"
	"image"
	"os"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

)

// Load the picture from the path

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open("C:/OnGithubAde/GoLang/Exercises/PixelGoLibrary")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// Convert the image to picturedata
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	cfg := pixelgl.WindowConfig{

			Title: "TestWindow",
			Bounds: pixel.R(0, 0, 284, 512),
			VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("background.png")
	if err != nil {
		panic(err)
	}

	background := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Aqua)

	background.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}

}


func main () {
	pixelgl.Run(run)
}

