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
	doneChannel := make(chan bool, 1)
	frameDoneChannel := make(chan bool, 1)
	frameRate := 24

	brightnessStrategy := image_to_ascii.AvarageBrightnessStrategyConcurrent{}
	converter := image_to_ascii.NewAsciiConverter(&brightnessStrategy)
	downloader := yt_video_downloader.NewYtDipVideoDownloader("144")
	downloadOut, err := downloader.DownloadToStdout("https://www.youtube.com/watch?v=QC8iQqtG0hg")

	renderer := renderer.NewTerminalRenderer(frameRate, framesChannel, doneChannel, frameDoneChannel)
	defer downloadOut.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	frameProcessor := video_to_ascii.NewVideoToFrameProcessor("pipe:0", "192x144", frameRate, converter, &downloadOut, framesChannel, frameDoneChannel)

	go frameProcessor.Process()
	//for buffering
	time.Sleep(5 * time.Second)
	go renderer.Render()
	//TODO should make a way to end
	<-doneChannel

}
