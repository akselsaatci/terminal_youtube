package image_to_ascii

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"runtime"
	"sync"
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

type AvarageBrightnessStrategyConcurrent struct {
}

func (a *AvarageBrightnessStrategyConcurrent) CalculateBrightness(img image.Image) ([][]uint32, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	res := make([][]uint32, width)
	for i := range res {
		res[i] = make([]uint32, height)
	}

	numCPU := runtime.NumCPU()
	numWorkers := min(numCPU, min(width, 32))

	if width*height < 100_000 {
		numWorkers = max(1, numWorkers/2)
	}
	type workItem struct {
		startX, endX int
	}
	jobs := make(chan workItem)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				for x := job.startX; x < job.endX; x++ {
					for y := 0; y < height; y++ {
						r, g, b, _ := img.At(x, y).RGBA()
						avg := (r + g + b) / 3 >> 8
						res[x][y] = uint32(avg)
					}
				}
			}
		}()
	}

	chunkSize := width / numWorkers
	if chunkSize == 0 {
		chunkSize = 1
	}

	for startX := 0; startX < width; startX += chunkSize {
		endX := startX + chunkSize
		if endX > width {
			endX = width
		}
		jobs <- workItem{startX, endX}
	}

	close(jobs)
	wg.Wait()

	return res, nil
}
