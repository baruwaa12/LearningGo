// Creating a game window

package main

import (
	"image"
	"os"
	_ "image/png"
	"fmt"

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
			// VSync: true,
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		spritesheet, err := loadPicture("trees.png")
		if err != nil {
			panic(err)
		}

		// NewBatch stores all the data of all currently drawn sprites so they dont have to be constantly redrawn after each
		// OpenGL draw calls
		batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)


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
			
		)

		var (
				frames = 0
				second = time.Tick(time.Second)
		)


		// last represents the current time
		// Duration is the difference between the last registered event and current time
		last := time.Now()
		for !win.Closed() {
				duration := time.Since(last).Seconds()
				last = time.Now()

				
			cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
			win.SetMatrix(cam)

			// These if statements show how much the camera should move given the length of a button press
			// When the left mouse button is pressed a tree should be drawn to the screen.
			if win.Pressed(pixelgl.MouseButtonLeft) {
					tree := pixel.NewSprite(spritesheet, treeFrames[rand.Intn(len(treeFrames))])
					mouse := cam.Unproject(win.MousePosition())
					tree.Draw(batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
			}
			if win.Pressed(pixelgl.KeyLeft) {
				camPos.X -= camSpeed * duration
			}
			if win.Pressed(pixelgl.KeyRight) {
				camPos.X += camSpeed * duration
			}
			if win.Pressed(pixelgl.KeyDown) {
				camPos.Y -= camSpeed * duration
			}
			if win.Pressed(pixelgl.KeyUp) {
				camPos.Y += camSpeed * duration
			}

			// Determines how far in the camera should move given a scroll
			camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

			win.Clear(colornames.Forestgreen)
			batch.Draw(win)
			win.Update()

			frames++
			select {
			case <-second:
					win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
					frames = 0
			default:
			}
		}
}


func main() {
		pixelgl.Run(run)
}
