// Creating a game window

package main

import (
	"image"
	"os"
	_ "image/png"

	"math"
	"math/rand"
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

		spritesheet, err := loadPicture("trees.png")
		if err != nil {
			panic(err)
		}

		var treeFrames []pixel.Rect
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
				for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
					treeFrames = append(treeFrames, pixel.R(x, y, x+32, y+32))
				}
		}
		

		// These are all the camera variables
		var (
			camPos = pixel.ZV
			camSpeed = 500.0
			camZoom = 1.0
			camZoomSpeed = 1.2
			trees  []*pixel.Sprite
			matrices []pixel.Matrix
		)


		last := time.Now()
		for !win.Closed() {
				dt := time.Since(last).Seconds()
				last = time.Now()

				
			cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
			win.SetMatrix(cam)

			// These if statements show how much the camera should move given the length of a button press
			if win.JustPressed(pixelgl.MouseButtonLeft) {
					tree := pixel.NewSprite(spritesheet, treeFrames[rand.Intn(len(treeFrames))])
					trees = append(trees, tree)
					mouse := cam.Unproject(win.MousePosition())
					matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
			}
			if win.Pressed(pixelgl.KeyLeft) {
				camPos.X -= camSpeed * dt
			}
			if win.Pressed(pixelgl.KeyRight) {
				camPos.X += camSpeed * dt
			}
			if win.Pressed(pixelgl.KeyDown) {
				camPos.Y -= camSpeed * dt
			}
			if win.Pressed(pixelgl.KeyUp) {
				camPos.Y += camSpeed * dt
			}

			// Determines how far in the camera should move given a scroll
			camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

			win.Clear(colornames.Forestgreen)

			for i, tree := range trees {
				tree.Draw(win, matrices[i])
			}

			win.Update()
		}
}



func main() {
		pixelgl.Run(run)
}
