package image_to_ascii

import (
	"image"
	_ "image/jpeg"

	"os"
	"testing"
)

const benchImagePath = "/Users/akselsaatci/Developer/huffman_encoding/pkg/image_to_ascii/image.jpg"

func BenchmarkConcurrent(b *testing.B) {
	bs := AvarageBrightnessStrategyConcurrent{}
	conv := NewAsciiConverter(&bs)

	f, err := os.Open(benchImagePath)
	if err != nil {
		b.Fatalf("Couldnt open the image\nError :%s", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)

	if err != nil {
		b.Fatalf("Couldnt decode the image\nError : %s", err)
	}
	b.ResetTimer() // Ensure only the conversion is timed
	b.StartTimer()
	conv.Convert(img)
	b.StopTimer()

}

func BenchmarkNormal(b *testing.B) {
	bs := AvarageBrightnessStrategy{}
	conv := NewAsciiConverter(&bs)

	f, err := os.Open(benchImagePath)
	if err != nil {
		b.Fatalf("Couldnt open the image\nError :%s", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)

	if err != nil {
		b.Fatalf("Couldnt decode the image\nError : %s", err)
	}

	b.ResetTimer() // Ensure only the conversion is timed

	b.StartTimer()
	conv.Convert(img)
	b.StopTimer()

}
