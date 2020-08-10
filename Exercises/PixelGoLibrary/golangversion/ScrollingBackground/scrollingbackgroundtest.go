// Scrolling Background test

package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	// "golang.org/x/image/colornames"
)

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
		Bounds: pixel.R(0, 0, 600, 500),
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
	
	// pic2, err := loadPicture("wingup.png")
	// if err != nil {
	// 	panic(err)
	// }


	background  := pixel.NewSprite(pic1, pic1.Bounds())
	// bird := pixel.NewSprite(pic2, pic2.Bounds())

	// win.Clear(colornames.Greenyellow)

	background.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	// bird.Draw(win, pixel.IM.Moved(win.Bounds().Center()))



	
	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}