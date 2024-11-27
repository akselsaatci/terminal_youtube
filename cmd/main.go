package main

import (
	"log"
	"time"

	"github.com/akselsaatci/terminal_youtube/pkg/image_to_ascii"
	"github.com/akselsaatci/terminal_youtube/pkg/renderer"
	"github.com/akselsaatci/terminal_youtube/pkg/video_to_ascii"
	"github.com/akselsaatci/terminal_youtube/pkg/yt_video_downloader"
)

func main() {

	framesChannel := make(chan string, 144)
	doneChannel := make(chan bool)
	frameRate := 24

	brightnessStrategy := image_to_ascii.AvarageBrightnessStrategyConcurrent{}
	converter := image_to_ascii.NewAsciiConverter(&brightnessStrategy)
	downloader := yt_video_downloader.NewYtDipVideoDownloader("480")
	downloadOut, err := downloader.DownloadToStdout("https://www.youtube.com/watch?v=u7kdVe8q5zs")

	renderer := renderer.NewTerminalRenderer(frameRate, framesChannel)
	defer downloadOut.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	frameProcessor := video_to_ascii.NewVideoToFrameProcessor("pipe:0", "640x480", frameRate, converter, &downloadOut, framesChannel)

	go frameProcessor.Process()
	//for buffering
	time.Sleep(5 * time.Second)
	go renderer.Render()
	//TODO should make a way to end
	for {
		if <-doneChannel {
			break
		}
	}

	close(framesChannel)

}
