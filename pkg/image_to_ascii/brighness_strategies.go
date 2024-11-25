package image_to_ascii

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type BrightnessStrategy interface {
	CalculateBrightness(img image.Image) ([][]uint32, error)
}

type AvarageBrightnessStrategy struct {
}

func (a *AvarageBrightnessStrategy) CalculateBrightness(img image.Image) ([][]uint32, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	res := make([][]uint32, width)
	for i := range res {
		res[i] = make([]uint32, height)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			avg := (r + g + b) / 3 >> 8
			res[x][y] = uint32(avg)
		}
	}

	return res, nil
}
