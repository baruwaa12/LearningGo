// Animation of the bird

package main

import (
	"os"
	"image"
	
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/errors"
	"golang.org/x/image/colornames"

)

// Load the bird animations. 
func loadBirdAnimations(sheetpath, descpath string, frameWidth float64) (sheet pixel.Picture, anims map[string][]pixel.Rect, err error) {

	defer func() {
		    if err != nil {
					err = errors.Wrap(err, "error loading animation sheet")
			}
	}()

	sheetFile, err := os.Open(sheetpath)
	if err != nil {
			return nil, nil, err
	}
	defer sheetFile.Close()
	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		return nil, nil, err
	}
	sheet = pixel.PictureDataFromImage(sheetImg)

	var frames []pixel.Rect
	for x := 0.0; x+frameWidth <= sheet.Bounds().Max.X; x += frameWidth {
			frames = append(frames, pixel.R(
				x,
				0,
				x+frameWidth,
				sheet.Bounds().H(),
			))
	}
	descFile, err := os.Open(descPath)
	if err != nil {
		return nil, nil, err
	}
	defer descFile.Close()

	anims = make(map[string][]pixel.Rect)

	// load the animation information, name and interval inside the spritesheet
	desc := csv.NewReader(descFile)
	for {
		anim, err := desc.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		name := anim[0]
		start, _ := strconv.Atoi(anim[1])
		end, _ := strconv.Atoi(anim[2])

		anims[name] = frames[start : end+1]
	}

	return sheet, anims, nil
}

}