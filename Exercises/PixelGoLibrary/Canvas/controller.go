package main

import (
	"github.com/faiface/pixel/pixelgl"
)

// controller handle input from devices
type controller struct {
	quit         bool
	moveLeftKey  pixelgl.Button
	moveRightKey pixelgl.Button
	moveClimbKey pixelgl.Button
	moveJumpKey  pixelgl.Button
	lightDt      float64
}

func (c *controller) create() {
	c.quit = false
}

func (c *controller) update() {
	if global.gWin.Pressed(pixelgl.KeyEscape){
		c.quit = true
	}

	if global.gActiveMenu != nil {
		if global.gWin.JustPressed(pixelgl.KeyUp) {
			global.gActiveMenu.moveUp()
		}
		if global.gWin.JustPressed(pixelgl.KeyDown) {
			global.gActiveMenu.moveDown()
		}
		if global.gWin.JustPressed(pixelgl.KeyEnter) {
			global.gActiveMenu.selectItem()
		}

	}
}