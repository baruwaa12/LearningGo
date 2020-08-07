package main

import (
	"image"
	"os"
	_ "image/png"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"time"
)


// Bird structure which contains all its properties
type bird struct {
	sprite  *pixel.Sprite
	sprite2 *pixel.Sprite
	frame   pixel.Rect
	flapup  bool
	lastFlapped time.Time

}


func (ba *bird) draw( win *pixelgl.Window)  {	

	// If the duration is up to 100 milliseconds it should switch images
	dt := time.Since(ba.lastFlapped).Milliseconds()
	if dt > 100 {	
		if ba.flapup {
			ba.sprite2.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		}else{
			ba.sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		}

		ba.lastFlapped = time.Now()
		ba.flapup = !ba.flapup

	}
}

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
	pic2, err := loadPictureBird("wingdown.png")
	if err != nil {
		panic(err)
	}

	// Load the picture as a sprite
	sprite := pixel.NewSprite(pic, pic.Bounds())
	sprite2 := pixel.NewSprite(pic2, pic2.Bounds())

	// // Draw the sprite to the center of the window
	// sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	// sprite2.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	 
	flappy := bird{sprite: sprite, sprite2: sprite2}

		

	for !win.Closed() {
		flappy.draw(win)
		win.Update()
	}

}

func main () {
	pixelgl.Run(run)
}
