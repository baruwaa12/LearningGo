// Configuration for the game. Loaded both from config file
// and static configuration.
package main

import (
	"encoding/json"
	"os"
	"github.com/faiface/pixel/pixelgl"

//=============================================================
// Game constants
//=============================================================
// General constants

)
const (
	GameTitle   = "Gizmo"
	GameVersion = "0.1"
)

var gSoundFiles = []string{
	"jump.wav",
	"shot.mp3",
	"shot2.mp3",
	"explosion.mp3",
	"shot3.mp3",
}

// World constants
const (
	wMaxInvFPS           = 1 / 60.0
	wShadowLength        = 5
	wShadowDepth         = 1.5
	wPixelsPerChunk      = 64
	wPixelSize           = 1
	wBorderSize          = 4
	wStaticBorderSize    = 0
	wStaticColor32       = 0xFFFFFFFE
	wStaticColor8        = 0xFE
	wFloodFill8          = 0xFC
	wFloodFill32         = 0xFFFFFFFC
	wFloodFillVisited8   = 0xFB
	wFloodFillVisited32  = 0xFFFFFFFB
	wBackground32        = 0xFFFFFF8F
	wBackground8         = 0x8F
	wBackgroundNew32     = 0xFFFFFF9F // New background when updated (so we don't have to update bg sprite)
	wBackgroundNew8      = 0x9F
	wLadder8             = 0xAF
	wLadder32            = 0xFFFFFFAF
	wShadow8             = 0xBF
	wShadow32            = 0xFFFFFFBF
	wViewMax             = 400 // 450
	wParticleDefaultLife = 5
	wGravity             = -9.82
	wAmmoMax             = 1000
	wDoorLen             = 30
	wDoorHeight          = 40
	wLightsMax           = 10
	wMiddleTextSize      = 22
	wFPSTextSize         = 40
	wDeathScreenText     = "You Died"
	wAssetObjectsPath    = "assets/objects/"
	wAssetMobsPath       = "assets/mobs/"
	wAssetMapsPath       = "assets/maps/"
	wAssetMixedPath      = "mixed/"
	wConfigFile          = "configuration.json"
)

// Global variables struct used throughtout the game.
type Global struct {
	gFont           *font
	gVariableConfig *variableConfig
	uTime           float32
	gWin			*pixelgl.Window
	gMenu           *menu
	gMenuItem       *menuItem
	gActiveMenu     *menu
	gControllerMenu *menu
	gController     *controller
}

var global = &Global{
	gVariableConfig: &variableConfig{},
	gFont:           &font{},
	gWin: 			 &pixelgl.Window{},
	gMenu:			 &menu{},
	gActiveMenu:     &menu{},
	gController:	 &controller{},
}

//=============================================================
// Variable configuration file that are possible to configure
// by the player from the menues.
//=============================================================
type variableConfig struct {
	Vsync             bool `json:"Vsync"`
	Fullscreen        bool `json:"Fullscreen"`
	WindowHeight      int  `json:"WindowHeight"`
	WindowWidth       int  `json:"WindowWidth"`
	UndecoratedWindow bool `json:"UndecoratedWindow"`
	CurrentMap        int  `json:"CurrentMap"`
	KeyShoot          int  `json:"KeyShoot"`
	KeyJump           int  `json:"KeyJump"`
	KeyLeft           int  `json:"KeyLeft"`
	KeyRight          int  `json:"KeyRight"`
	KeyClimb          int  `json:"KeyClimb"`
	KeyAction         int  `json:"KeyAction"`
	KeyDrop           int  `json:"KeyDrop"`
	KeyPickup         int  `json:"KeyPickup"`
	KeyDuck           int  `json:"KeyDuck"`
	MaxParticles      int  `json:"MaxParticles"`
}

// loadConfiguration Configuration file loading
func (v *variableConfig) LoadConfiguration() {
	configFile, err := os.Open(wConfigFile)
	defer configFile.Close()
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(v)
	if err != nil {
		panic("Configuration file not valid.")
	}
}

