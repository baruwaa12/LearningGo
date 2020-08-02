// Creating a game window

package main

import (
	"image"
	"os"
	_ "image/png"

	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)


// This function loads pictures from a path within the system.
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

	// This func converts the image to picturedata
	return pixel.PictureDataFromImage(img), nil
}

func run() {
		cfg := pixelgl.WindowConfig{

			// Defines Window properties
			Title: "Pixel rocks",
			Bounds: pixel.R(0, 0, 1024, 728),
			VSync: true,
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}


		// Smooths out the image each time it is drawn
		win.SetSmooth(true)

		pic, err := loadPicture("birdwingup.png")
		if err != nil {
			    panic(err)
		}

		sprite := pixel.NewSprite(pic, pic.Bounds())

		angle := 0.0

		last := time.Now()
		for !win.Closed() {
				dt := time.Since(last).Seconds()
				last = time.Now()

				angle += 3 * dt

				win.Clear(colornames.Firebrick)


				
				mat := pixel.IM

				// This will rotate the image by a number of radians around a given vector 
				mat = mat.Rotated(pixel.ZV, angle)
				mat = mat.Moved(win.Bounds().Center())
				sprite.Draw(win, mat)

				win.Update()
		}
}



func main() {
		pixelgl.Run(run)
}
