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
	x 			int
	y			int
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


// Recenter the birds position after the bird collides with a pipe
func (ba *bird) recenterPosition(win *pixelgl.Window) {
	ba.y = win.Bounds().Center().Y
}

func (ring *Ring) draw(win *pixelgl.Window) {

	
	
	startingpostion := rand.Intn((900 - 100) + 100)
    currentVec := pixel.Vec{X: float64(startingpostion), Y:float64(startingpostion)}


	ring.sprite.Draw(global.gWin, pixel.IM.Moved(currentVec))
}

// Extend PipePair type by add draw function.
// Input: Window = the game window that will be drawn on.
// Draws pipe pair scrolling across the screen.

func (pipes *PipePair) draw(win *pixelgl.Window) {


	// Grab secomds since we last drew the pipes
	dt := time.Since(pipes.lastDrawn).Milliseconds();
	
	// If the duration is over 10 milliseconds 
	// Pipe moves across the screen 1 pixel.
	if dt >= 10 {
		pipes.x = pipes.x - 1
		pipes.lastDrawn = time.Now()
	
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

	startingpostion := rand.Intn((900 - 100) + 100)

	// New instances of the pipepair using PipePairs struct
	pipepair := PipePair{FacingUp: pipeUpPic, FacingDown: pipeDownPic, x: x/2, yUp: rectUp.Y, yDown: rectDown.Y,  lastDrawn: time.Now()}
	pipepair2 := PipePair{FacingUp: pipeUpPic, FacingDown: pipeDownPic, x: x, yUp: rectUp.Y, yDown: rectDown.Y,  lastDrawn: time.Now()}
	flappy := bird{sprite: sprite, sprite2: sprite2, x: win.Bounds().Center().X, y :win.Bounds().Center().Y }
	ringobject := Ring{sprite: ring, x: startingpostion, y: startingpostion}
	global.gGameScene = &GameScene{background: background, flappy: flappy, obstacle1: pipepair, obstacle2: pipepair2, Ring: ringobject}
}



func drawGameScene() {
	var win = global.gWin

	global.gGameScene.Ring.draw(win)
	global.gGameScene.background.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	global.gGameScene.flappy.draw(win)
	global.gGameScene.obstacle1.draw(win)
	global.gGameScene.obstacle2.draw(win)
	win.Update()
}

