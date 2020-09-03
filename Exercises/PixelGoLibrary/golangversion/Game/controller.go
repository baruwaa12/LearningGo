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
	if global.gWin.Pressed(pixelgl.KeyQ){
		c.quit = true
	}

	if global.gActiveMenu != nil {

		// if the mouse hovers within the area of the buttons it should highlight that specfic menu item


			mouseBounds := global.gWin.MousePosition()
	

			for i, item := range global.gActiveMenu.items {
				itemSize := item.canvas.Bounds()
				itemBounds1 := item.position

				minX := itemBounds1.X - (itemSize.Max.X*0.5)
				maxX := itemBounds1.X + (itemSize.Max.X*0.5)
	
				minY := itemBounds1.Y - (itemSize.Max.Y*0.5)
				maxY := itemBounds1.Y + (itemSize.Max.Y*0.5)
				
				// If the mouse is within the bounds of a button it should highlight the button
				if mouseBounds.Y <= maxY && mouseBounds.Y >= minY  && 
						mouseBounds.X <= maxX && mouseBounds.X >= minX {
							// 
					// Select current menu
					global.gActiveMenu.items[i].selected = 1
						
					if global.gWin.JustPressed(pixelgl.MouseButtonLeft) {
						global.gActiveMenu.selectItem()
					}
					// unselects menu items
					if global.gActiveMenu != nil {
						for z := range global.gActiveMenu.items {
							if i != z {
								global.gActiveMenu.items[z].selected = 0
							}	

						}
					}
				}
			}

		
		if global.gWin.JustPressed(pixelgl.KeyUp) {
			global.gActiveMenu.moveUp()
		}
		if global.gWin.JustPressed(pixelgl.KeyDown) {
			global.gActiveMenu.moveDown()
		}
		if global.gWin.JustPressed(pixelgl.KeyEnter) {
			global.gActiveMenu.selectItem()
		}

		// Terminate game when menu is showing 
		// and user clicks escape key.
		if global.gWin.Pressed(pixelgl.KeyEscape) {
			c.quit = true
		}
		
	}else {
		if global.gWin.JustPressed(pixelgl.KeyUp) {
			global.gGameScene.flappy.y = global.gGameScene.flappy.y + 20
		}
		if global.gWin.JustPressed(pixelgl.KeyEscape) {
			global.gActiveMenu = global.gMenu
		}
		return
	}

}