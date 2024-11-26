package image_to_ascii

import (
	"image"
	"math"
	"runtime"
	"strings"
	"sync"
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

func (c *AsciiConverter) ConvertConcurrent(img image.Image) (string, error) {
	brightnessLevels, err := c.brightnessStrategy.CalculateBrightness(img)
	if err != nil {
		return "", err
	}

	return c.generateAsciiConcurrent(brightnessLevels), nil
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

func (c *AsciiConverter) generateAsciiConcurrent(input [][]uint32) string {
	if len(input) == 0 || len(input[0]) == 0 {
		return ""
	}

	height := len(input[0])
	divAmount := 255.0 / float64(len(c.charset))

	numCPU := runtime.NumCPU()
	numWorkers := min(numCPU, min(height, 32))
	if height < 100 {
		numWorkers = 1
	}

	type workItem struct {
		startY, endY int
	}
	jobs := make(chan workItem)
	results := make([]string, height)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var builder strings.Builder
			for job := range jobs {
				for y := job.startY; y < job.endY; y++ {
					builder.Reset()
					for x := 0; x < len(input); x++ {
						charIndex := int(math.Floor(float64(input[x][y]) / divAmount))
						if charIndex >= len(c.charset) {
							charIndex = len(c.charset) - 1
						}
						builder.WriteByte(c.charset[charIndex])
					}
					results[y] = builder.String()
				}
			}
		}()
	}

	chunkSize := height / numWorkers
	if chunkSize == 0 {
		chunkSize = 1
	}

	for startY := 0; startY < height; startY += chunkSize {
		endY := startY + chunkSize
		if endY > height {
			endY = height
		}
		jobs <- workItem{startY, endY}
	}

	close(jobs)
	wg.Wait()

	var finalBuilder strings.Builder
	finalBuilder.Grow(height * (len(input) + 1))
	for y := 0; y < height; y++ {
		finalBuilder.WriteString(results[y])
		finalBuilder.WriteByte('\n')
	}

	return finalBuilder.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
