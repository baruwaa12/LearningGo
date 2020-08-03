package main

import (
	"image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var gopherimg *pixel.Sprite

func gameloop(win *pixelgl.Window) {
	win.Canvas().SetFragmentShader(fragmentShader)

	for !win.Closed() {
			win.Update()
	}
}

var fragmentShader = `
#version 330 core
// The first line in glsl source code must always start with a version directive as seen above.

// vTexCoords are the texture coordinates, provided by Pixel
in vec2  vTexCoords;

// fragColor is what the final result is and will be rendered to your screen.
out vec4 fragColor;

// uTexBounds is the texture's boundries, provided by Pixel.
uniform vec4 uTexBounds;

// uTexture is the actualy texture we are sampling from, also provided by Pixel.
uniform sampler2D uTexture;

void main() {
		// t is set to the screen coordinate for the current fragment.
		vec2 t = (vTexCoords - uTexBounds.xy) / uTexBounds.zw;

		// Sum our 3 color channels
	  float sum  = texture(uTexture, t).r;
			sum += texture(uTexture, t).g;
			sum += texture(uTexture, t).b;

		// Divide by 3, and set the output to the result
		vec4 color = vec4( sum/3, sum/3, sum/3, 1.0);
		fragColor = color;
		  
}
`


func run() {
		cfg := pixelgl.WindowConfig{
				Title: "Pixel Rocks!",
				Bounds:  pixel.R(0, 0, 325, 348),
				VSync: true,
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}
		f, err := os.Open("C:/OnGithubAde/GoLang/Exercises/PixelGoLibrary/thegopherproject.png")
		if err != nil {
			panic(err)
		}
		img, err := png.Decode(f)
		if err != nil {
			panic(err)
		}
		pd := pixel.PictureDataFromImage(img)
		gopherimg = pixel.NewSprite(pd, pd.Bounds())
		gopherimg.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		gameloop(win)
}


func main() {
	pixelgl.Run(run)
}