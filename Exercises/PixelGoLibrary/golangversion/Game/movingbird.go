package main

import (
	"fmt"
	"time"
	"image"
	"os"

	_ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
)

// PipePair structure which contains all its properties
type PipePair struct {
	FacingUp   	pixel.Picture
	FacingDown	pixel.Picture	
	x			float64
	yUp 		float64
	yDown 		float64
	lastDrawn	time.Time
}


type Ring struct {
	sprite  	*pixel.Sprite
	x 			float64
	y			float64
	lastDrawn	time.Time

}

type bird struct {
	sprite  *pixel.Sprite
	sprite2 *pixel.Sprite
	frame   pixel.Rect
	flapup  bool
	lastFlapped time.Time
	y       float64
	x		float64
}

type GameScene struct {
	background  *pixel.Sprite
	flappy  	bird
	obstacle1	PipePair
	obstacle2	PipePair
	Ring		Ring
}

func (ba *bird) draw( win *pixelgl.Window)  {	

	// If the duration is up to 100 milliseconds it should switch images

	currentVec := pixel.Vec{X: ba.x, Y:ba.y}

	dt := time.Since(ba.lastFlapped).Milliseconds()
	if dt > 100 {	
		ba.lastFlapped = time.Now()
		ba.flapup = !ba.flapup
	}

	if dt > 25 {
		ba.y = ba.y - 0.2
	}
	
	if ba.flapup {
		ba.sprite2.Draw(win, pixel.IM.Moved(currentVec))
	}else{
		ba.sprite.Draw(win, pixel.IM.Moved(currentVec))
	}
}

func randomfloat(min, max float64) float64 {
	position := min + max + rand.Float64()
	return position
}

func (ba *bird) CollidedWithPipe( pipes PipePair) bool  {	

	currentVec := pixel.Vec{X: ba.x, Y:ba.y}

	surfaceX := currentVec.X + 16
	surfaceY := currentVec.Y - 16
	surfaceTopY := currentVec.Y + 16

	pipeDownY := pipes.yDown + 75
	pipeDownX := pipes.x - 40
	pipeUpY   := pipes.yUp - 75
	
	if surfaceY <= pipeDownY && (surfaceX >= pipeDownX && surfaceX <= pipes.x) {
		//bottom  collision
		fmt.Println("PipeDown Collision")
		return true
	}

	if surfaceTopY >= pipeUpY && (surfaceX >= pipeDownX && surfaceX <= pipes.x) {
		// Top pipe collision
		fmt.Println("PipeUp collision")
		return true
	}

	return false
}

// func randFloats(min, max float64) float64 {
// 	for i := range result {
// 		result = min + rand.Float64() * (max - min)
// 	}
// 	return result
// }

// Recenter the birds position after the bird collides with a pipe
func (ba *bird) recenterPosition(win *pixelgl.Window) {
	ba.y = win.Bounds().Center().Y
}


func (ring *Ring) draw(win *pixelgl.Window) {

	currentVec1 := pixel.Vec{X: ring.x, Y: ring.y}
	// currentVec2 := pixel.Vec{X: ring.x, Y: ring.y + 50}


	// Grab seconds since we last drew the pipes
	dt := time.Since(ring.lastDrawn).Milliseconds();
	
	// If the duration is over 10 milliseconds 
	// Pipe moves across the screen 1 pixel.
	if dt >= 10 {
		ring.x = ring.x - 1 
		ring.lastDrawn = time.Now()
	}

	ring.y = ring.y
	
	if ring.x  < 20 {
		ring.x = 1100
		ring.y = ring.y - 600
		if ring.y < - 80 {
			ring.y = 500
		}
	}




	ring.sprite.Draw(global.gWin, pixel.IM.Moved(currentVec1))
}


func (ring2 *Ring) draw(win *pixelgl.Window) {

	currentVec := pixel.Vec{}
}


// Extend PipePair type by add draw function.
// Input: Window = the game window that will be drawn on.
// Draws pipe pair scrolling across the screen.

func (pipes *PipePair) draw(win *pixelgl.Window) {


	// Grab seconds since we last drew the pipes
	dt := time.Since(pipes.lastDrawn).Milliseconds();
	
	// If the duration is over 10 milliseconds 
	// Pipe moves across the screen 1 pixel.
	if dt >= 10 {
		pipes.x = pipes.x - 1
		pipes.lastDrawn = time.Now()
		if dt >= 5000 {
			pipes.x = pipes.x - 50
		}
	}

	// if dt >= 1000 {
	// 	pipes.x = pipes.x - 20
	// 	pipes.lastDrawn = time.Now()

	// }	

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

}

func gameSetup() {

	var win = global.gWin

	WingUp, err := loadPicture("wingup.png")
	if err != nil {
		panic(err)
	}
	WingDown, err := loadPicture("wingdown.png")
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

	ringpic, err := loadPicture("ring1.png")
	if err != nil {
		panic(err)
	}

	background  := pixel.NewSprite(pic1, pic1.Bounds())
	sprite := pixel.NewSprite(WingUp, WingUp.Bounds())
	sprite2 := pixel.NewSprite(WingDown, WingDown.Bounds())
	ring := pixel.NewSprite(ringpic, ringpic.Bounds())

	
	winRect := win.Bounds().Center()
	rectDown := pixel.Vec{winRect.X, win.Bounds().H()/4}
	rectUp := pixel.Vec{winRect.X, (win.Bounds().H()*0.75)}

	x := win.Bounds().W()

	// startingpostion := win.Bounds().Center()

	// New instances of the pipepair using PipePairs struct
	pipepair := PipePair{FacingUp: pipeUpPic, FacingDown: pipeDownPic, x: x/2, yUp: rectUp.Y, yDown: rectDown.Y,  lastDrawn: time.Now()}
	pipepair2 := PipePair{FacingUp: pipeUpPic, FacingDown: pipeDownPic, x: x, yUp: rectUp.Y, yDown: rectDown.Y,  lastDrawn: time.Now()}
	flappy := bird{sprite: sprite, sprite2: sprite2, x: win.Bounds().Center().X, y :win.Bounds().Center().Y }
	ringobject := Ring{sprite: ring, x: x, y: x}
	// ringobject2 := Ring{sprite: ring, x: x, y : (x - 100)}
	global.gGameScene = &GameScene{background: background, flappy: flappy, obstacle1: pipepair, obstacle2: pipepair2, Ring: ringobject}
}



func drawGameScene() {
	var win = global.gWin

	global.gGameScene.background.Draw(win, pixel.IM.Moved(win.Bounds().Center())) 
	global.gGameScene.flappy.draw(win)
	global.gGameScene.obstacle1.draw(win)
	global.gGameScene.obstacle2.draw(win)
	global.gGameScene.Ring.draw(win)


	win.Update()
}

// Need to fix placement and also reduce speed at which the object is being painted on screen.
// Multiple rings on the screen 
