package image_to_ascii

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

const charset string = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
const divAmount float32 = 3.54166

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
	for i := 0; i < imageW; i++ { // Outer slice
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
			r := uint8(cR >> 8)
			g := uint8(cG >> 8)
			b := uint8(cB >> 8)
			avg := (r + g + b) / 3

			res[i][j] = uint32(avg)
		}
	}

	return res, nil
}

func generateBrightnessToAscii(input [][]uint32) string {
	res := ""
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			charIndex := int(float32(input[i][j]) / divAmount)
			if charIndex >= len(charset) {
				charIndex = len(charset) - 1
			}

			res += string(charset[charIndex])
		}
		res += "\n"
	}
	return res
}

func ImageToAscii(imagePath string) (string, error) {
	img, conf, err := openImage(imagePath)

	if err != nil {
		return "", err
	}

	brightnessLevels, err := calculateBrightnessOfPixels(img, conf)

	if err != nil {
		return "", err
	}

	return generateBrightnessToAscii(brightnessLevels), nil
}
