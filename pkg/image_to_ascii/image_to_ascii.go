package image_to_ascii

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
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

const charset string = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "

type AsciiConverter struct {
	brightnessStrategy BrightnessStrategy
	charset            string
}

func NewAsciiConverter() *AsciiConverter {
	return &AsciiConverter{
		brightnessStrategy: &AvarageBrightnessStrategy{},

		charset: "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ",
	}
}

func (c *AsciiConverter) Convert(img image.Image) (string, error) {
	brightnessLevels, err := c.brightnessStrategy.CalculateBrightness(img)
	if err != nil {
		return "", err
	}

	return c.generateAscii(brightnessLevels), nil
}

func (c *AsciiConverter) generateAscii(input [][]uint32) string {
	if len(input) == 0 || len(input[0]) == 0 {
		return ""
	}

	divAmount := 255.0 / float64(len(c.charset))
	var res string

	for y := 0; y < len(input[0]); y++ {
		for x := 0; x < len(input); x++ {
			charIndex := int(math.Floor(float64(input[x][y]) / divAmount))
			if charIndex >= len(c.charset) {
				charIndex = len(c.charset) - 1
			}
			res += string(c.charset[charIndex])
		}
		res += "\n"
	}

	return res
}

func ImageToAscii(img image.Image) (string, error) {
	converter := NewAsciiConverter()
	return converter.Convert(img)
}
