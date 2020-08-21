package main

import (
	"time"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

)

// Main
func main() {
	//	defer profile.Start(profile.CPUProfile).Stop()
	pixelgl.Run(run)
}


type menu struct {
	items   []*menuItem
	logo    *menuItem
	canvas  *pixelgl.Canvas
}

type menuItem struct {
	canvas  *pixelgl.Canvas
	scale   float64
	name    string
	selected  int32
	action   func()

}

// Function to create the main menu 
func (m *menu)  CreateMainMenu(){

	// Properties of the menu logo
	m.logo = &menuItem{
		canvas:  pixelgl.NewCanvas(pixel.R(0, 0, 1, 1)),
		scale:   2.0,
		name:    "Flappy",
		selected: 2,
	}
	m.logo.canvas.SetUniform("uSelected", m.logo.selected)
	m.logo.canvas.SetUniform("uTime", &global.uTime)
	m.logo.canvas.SetFragmentShader(fragmentShaderMenuItem)

	// All Menu items are appended to the array of menu items
	// Use method addMenuItem to create multiple menu items
	m.items = make([]*menuItem, 0)
	m.addMenuItem(1.0, "Press screen to start", func() { fmt.Println("Screen Start")  })
	m.addMenuItem(1.0, "Continue", func() { fmt.Println("Screen Continue") })
	m.items[0].selected = 1
}

// Extend menu method to add menu items
func (m *menu) addMenuItem(menuItemScale float64, menuItemName string, menuItemAction func() ) {
	item := &menuItem{
		canvas:   pixelgl.NewCanvas(pixel.R(0, 0, 1, 1)),
		name:     menuItemName,
		action:   menuItemAction,
		selected: 0,
		scale:    menuItemScale,
	}
	item.canvas.SetUniform("uSelected", &item.selected)
	item.canvas.SetUniform("uTime", &global.uTime)
	item.canvas.SetFragmentShader(fragmentShaderMenuItem)

	m.items = append(m.items, item)
}

func (m *menu) moveUp() {
	for i, item := range m.items {
		if item.selected == 1 {
			if i > 0 {
				m.items[i-1].selected = 1
			} else {
				m.items[len(m.items) - 1].selected = 1
			}
			m.items[i].selected = 0
			return
		}

	}
}

func (m *menu) moveDown() {
	for i, item := range m.items {
		if item.selected == 1 {
			if i == len(m.items) - 1 {
				m.items[0].selected = 1
			} else {
				m.items[i + 1].selected = 1
			}
			m.items[i].selected = 0
			return
		}
	}
}

func (m *menu) selectItem() {
	for i, item := range m.items {
		if item.selected == 1 {
			m.items[i].action()
			return
		}
	}
}

func (m *menu) draw() {

	center := global.gWin.Bounds().Center()
	
	if m.logo != nil {
		m.logo.canvas.Clear(pixel.RGBA{R: 0, G: 0, B: 0, A: 0})	
		global.gFont.writeToCanvas("Flappy", m.logo.canvas)
		// offsetX := (float64(global.gVariableConfig.WindowWidth) / 2) / 2 //- m.logo.canvas.Bounds().Max.X/2
		offsetY := (float64(global.gVariableConfig.WindowHeight) / 3) - m.logo.canvas.Bounds().Max.Y
		logoScale := float64(global.gVariableConfig.WindowWidth * 100) / float64(global.gVariableConfig.WindowHeight * 30)
		m.logo.canvas.Draw(global.gWin, pixel.IM.Scaled(pixel.ZV, logoScale).Moved(pixel.V(center.X, center.Y+offsetY)))
	}
	for i := range m.items {
		m.items[i].canvas.Clear(pixel.RGBA{R: 0, G: 0, B: 0, A: 0})	
		global.gFont.writeToCanvas(m.items[i].name, m.items[i].canvas)
		itemScale := (float64(global.gVariableConfig.WindowWidth) / float64(global.gVariableConfig.WindowHeight)) / 1
		// offsetX := (float64(global.gVariableConfig.WindowWidth) / 3) / 2 //- m.logo.canvas.Bounds().Max.X/2
		offsetY := (float64(global.gVariableConfig.WindowHeight) / 1.6 / 3) - m.items[i].canvas.Bounds().Max.Y*float64(i)*itemScale
		m.items[i].canvas.Draw(global.gWin, pixel.IM.Scaled(pixel.ZV, itemScale).Moved(pixel.V(center.X, center.Y-offsetY)))
	}
}

func run() {

	global.gVariableConfig.LoadConfiguration()

	cfg := pixelgl.WindowConfig{
		Title:       "PacMan",
		Bounds:      pixel.R(0, 0, float64(1080), float64(768)),
		VSync:       false,
		Undecorated: false,
	}
	
	gWin, err := pixelgl.NewWindow(cfg)

	Debug("W:", gWin.Bounds())

	gWin.SetCursorVisible(true)

	if err != nil {
		panic(err)
	}

	centerWindow(gWin)

	global.gWin = gWin

	setup()

	gameLoop()

}

func setup() {
	global.gFont.create()

	global.gMenu.CreateMainMenu()

	global.gActiveMenu = global.gMenu

	global.gController.create()

	global.gWin.SetSmooth(false)

	global.gWin.Canvas().SetFragmentShader(fragmentShaderFullScreen)
}

func gameLoop() {

	last := time.Now()
	frameDt := 0.0

	
	for !global.gWin.Closed() {
		dt := time.Since(last).Seconds()
		frameDt += dt
		last = time.Now()

		if global.gActiveMenu != nil {
			global.gActiveMenu.draw()
		}
		global.gController.update()

		global.gWin.Update()
	}		
}