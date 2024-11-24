package main

import (
	"fmt"
	"image/png"
	"io"
	"os/exec"

	"github.com/akselsaatci/huffman/pkg/image_to_ascii"
)

func main() {
	// FFmpeg command to extract frames as PNG and write to stdout
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i", "/Users/akselsaatci/Developer/huffman_encoding/cmd/file.mp4",
		"-vf", "fps=6",
		"-vsync", "vfr",
		"-f", "image2pipe",
		"-vcodec", "png",
		"-s", "640x480",
		"-",
	)

	stdout, err := ffmpegCmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating stdout pipe:", err)
		return
	}

	if err := ffmpegCmd.Start(); err != nil {
		fmt.Println("Error starting FFmpeg command:", err)
		return
	}

	for {
		// Decode PNG frame from stdout
		img, err := png.Decode(stdout)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding PNG:", err)
			break
		}

		fmt.Printf("Decoded frame: %dx%d\n", img.Bounds().Dx(), img.Bounds().Dy())

		if img != nil {
			fmt.Print("\033[H\033[2J")
			fmt.Println(image_to_ascii.ImageToAscii(img))
		}

	}

	if err := ffmpegCmd.Wait(); err != nil {
		fmt.Println("Error waiting for FFmpeg command:", err)
		return
	}

	fmt.Println("Frames extracted and processed.")
}
