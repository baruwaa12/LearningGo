 package main

import (
	"log"
	"os"
	"golang.org/x/image/colornames"
	"time"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep"
	
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
	canvas        *pixelgl.Canvas
	scale         float64
	name          string
	selected      int32
	action    	  func()
	position	  pixel.Vec
	itemBounds	  pixel.Vec

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
	m.addMenuItem(1.0, "Quit", func() {
		fmt.Println("Screen to quit")
		global.gActiveMenu = nil
		global.gController.quit = true

	 })
	m.addMenuItem(1.0, "Press enter to start", func() { 
		fmt.Println("Screen Continue")
		global.gActiveMenu = nil
		global.gGameScene.flappy.recenterPosition(global.gWin) })
	m.addMenuItem(1.0, "Options", func() {
		global.gActiveMenu = global.gOptionsMenu
		fmt.Println("Hello")
		
	})
	m.items[0].selected = 1 
}

func (m *menu) createOptions() {
	m.CreateMainMenu()

	m.addMenuItem(0.7, "Controls", func() {
	})
	m.addMenuItem(0.8, "Back", func() {
		global.gActiveMenu = global.gMenu
	})
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
	fmt.Println(center)
	
	if m.logo != nil {
		m.logo.canvas.Clear(pixel.RGBA{R: 0, G: 0, B: 0, A: 0})	
		global.gFont.writeToCanvas("Flappy", m.logo.canvas)
		offsetY := (float64(global.gVariableConfig.WindowHeight) / 3) - m.logo.canvas.Bounds().Max.Y
		logoScale := float64(global.gVariableConfig.WindowWidth * 100) / float64(global.gVariableConfig.WindowHeight * 30)
		m.logo.canvas.Draw(global.gWin, pixel.IM.Scaled(pixel.ZV, logoScale).Moved(pixel.V(center.X, center.Y+offsetY)))
	}
	for i := range m.items {
		m.items[i].canvas.Clear(pixel.RGBA{R: 0, G: 0, B: 0, A: 0})	
		global.gFont.writeToCanvas(m.items[i].name, m.items[i].canvas)
		itemScale := (float64(global.gVariableConfig.WindowWidth) / float64(global.gVariableConfig.WindowHeight)) / 1
		offsetY := (float64(global.gVariableConfig.WindowHeight) / 1.6 / 3) - m.items[i].canvas.Bounds().Max.Y*float64(i)*itemScale
		m.items[i].canvas.Draw(global.gWin, pixel.IM.Scaled(pixel.ZV, itemScale).Moved(pixel.V(center.X, center.Y-offsetY)))
		m.items[i].position = pixel.Vec{X: center.X, Y: center.Y-offsetY}
	}
}


func menuSetup() {
	
	global.gOptionsMenu.createOptions()

	global.gFont.create()

	global.gMenu.CreateMainMenu()

	global.gActiveMenu = global.gMenu

	global.gController.create()

	global.gWin.SetSmooth(false)

	
	global.gWin.Canvas().SetFragmentShader(fragmentShaderFullScreen)
}

func run() {

	global.gVariableConfig.LoadConfiguration()

	cfg := pixelgl.WindowConfig{
		Title:       "Flappy",
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

	menuSetup()

	gameSetup()

	gameLoop()

}

func musicSetup() beep.StreamSeekCloser{
	f, err := os.Open("Observatory.mp3")
	if err != nil {
			log.Fatal(err)
		}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	
	speaker.Play(streamer)

	return streamer
}



func gameLoop() {

	last := time.Now()
	frameDt := 0.0

	streamer := musicSetup()
	

	for !global.gWin.Closed() && !global.gController.quit {
	
		dt := time.Since(last).Seconds()
		frameDt += dt
		last = time.Now()


		global.gWin.Clear(global.gClearColor)

		if global.gActiveMenu != nil {
			global.gActiveMenu.draw()
			global.gWin.Update()
		} else{
			

			drawGameScene()
			if global.gGameScene.flappy.CollidedWithPipe(global.gGameScene.obstacle1)  || global.gGameScene.flappy.CollidedWithPipe(global.gGameScene.obstacle2) {
				global.gWin.Clear(colornames.Black)
				global.gActiveMenu = global.gMenu
				global.gWin.Update()
			}
		}
		
		global.gController.update()

		global.gWin.Update()
	}
	streamer.Close()	

}

