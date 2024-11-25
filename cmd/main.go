package main

import (
	"log"

	"github.com/akselsaatci/terminal_youtube/pkg/image_to_ascii"
	"github.com/akselsaatci/terminal_youtube/pkg/video_to_ascii"
	"github.com/akselsaatci/terminal_youtube/pkg/yt_video_downloader"
)

func main() {
	brightnessStrategy := image_to_ascii.AvarageBrightnessStrategy{}
	converter := image_to_ascii.NewAsciiConverter(&brightnessStrategy)

	downloader := yt_video_downloader.NewYtDipVideoDownloader("144")

	downloadOut, err := downloader.DownloadToStdout("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer downloadOut.Close()

	frameProcessor := video_to_ascii.NewVideoToFrameProcessor("pipe:0", "192x144", 24, *converter, &downloadOut)
	frameProcessor.Process()
}
