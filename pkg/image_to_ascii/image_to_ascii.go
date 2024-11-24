package image_to_ascii

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

const charset string = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "

func openImage(filePath string) (image.Image, *image.Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)

	if err != nil {
		panic("Couldnt decode image")
	}

	_, err = f.Seek(0, 0)

	config, _, err := image.DecodeConfig(f)

	return img, &config, err
}

func calculateBrightnessOfPixels(img image.Image, config *image.Config) ([][]uint32, error) {
	imageW := config.Width
	imageH := config.Height

	res := make([][]uint32, imageW)
	for i := 0; i < imageW; i++ {
		res[i] = make([]uint32, imageH)
	}

	if img == nil {
		return nil, errors.New("Image is null!")
	}
	if config == nil {
		return nil, errors.New("Config is null!")
	}

	for i := 0; i < imageW; i++ {
		for j := 0; j < imageH; j++ {
			currentColor := img.At(i, j)
			cR, cG, cB, _ := currentColor.RGBA()
			r := uint32(cR >> 8)
			g := uint32(cG >> 8)
			b := uint32(cB >> 8)
			avg := (r + g + b) / 3

			res[i][j] = uint32(avg)
		}
	}

	return res, nil
}

func generateBrightnessToAscii(input [][]uint32) (string, error) {
	// does this necessary :?
	if len(input[0]) <= 0 {
		return "", errors.New("Input Should have elements")
	}

	divAmount := 255.0 / float64(len(charset))
	res := ""
	for j := 0; j < len(input[0]); j++ {
		for i := 0; i < len(input); i++ {
			charIndex := int(math.Floor(float64(input[i][j])/divAmount)) + 1
			if charIndex >= len(charset) {
				charIndex = len(charset) - 1
			}
			res += string(charset[charIndex])
		}
		res += "\n"
	}
	return res, nil
}

func ImageToAscii(img image.Image) (string, error) {
	conf := image.Config{Width: 640, Height: 480}

	brightnessLevels, err := calculateBrightnessOfPixels(img, &conf)

	if err != nil {
		return "", err
	}
	res, err := generateBrightnessToAscii(brightnessLevels)
	if err != nil {
		return "", err
	}

	return res, nil
}
