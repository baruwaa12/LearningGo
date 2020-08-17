package main

import (
	"image"
	"os"

	_ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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

	buttonBackground := pixel.NewSprite(pic1, pic1.Bounds())

	// While the window is not closed, after every refresh the pipepairs should be drawn at its next position.
	for !win.Closed() {	
		createCanvas(win, buttonBackground)
		win.Update()
		
	}
}


func createCanvas(win *pixelgl.Window, sprite *pixel.Sprite) string  {
	canvas:= pixelgl.NewCanvas(pixel.R(100, 100, 100, 100))
	canvas.Clear(pixel.RGBA{R: 0, G: 0, B: 0, A: 0})
	global.gFont.writeToCanvas("Menu", canvas)
	// canvas.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	itemScale := (float64(win.Bounds().W()) / float64(win.Bounds().H())) / 3
	offsetX := (win.Bounds().W() / 1) / 2 
	offsetY := (float64(win.Bounds().H()) / 1.6 / 1) - canvas.Bounds().Max.Y * float64(1) * itemScale

	canvas.Draw(win, pixel.IM.Scaled(pixel.ZV, itemScale).Moved(pixel.V(1 + offsetX, 1 + offsetY)))
	win.Update()

	canvas2:= pixelgl.NewCanvas(pixel.R(100, 100, 100, 100))
	canvas2.Clear(pixel.RGBA{R: 0, G: 0, B: 0, A: 0})
	global.gFont.writeToCanvas("Menu 2", canvas2)
	// canvas.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	offsetY2 := (float64(win.Bounds().H()) / 1.6 / 1) - canvas.Bounds().Max.Y * float64(3) * itemScale

	canvas2.Draw(win, pixel.IM.Scaled(pixel.ZV, itemScale).Moved(pixel.V(1 + offsetX, 1 + offsetY2)))
	win.Update()

	return "done"
}


func main() {
	pixelgl.Run(run)
	
}