package image_to_ascii

import (
	"image"
	"math"
)

type AsciiConverter struct {
	brightnessStrategy BrightnessStrategy
	charset            string
}

func NewAsciiConverter(b BrightnessStrategy) *AsciiConverter {
	return &AsciiConverter{
		brightnessStrategy: b,

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
