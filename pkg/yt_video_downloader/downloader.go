package yt_video_downloader

import (
	"io"
	"os/exec"
)

type YoutubeVideoDownloader interface {
	DownloadToStdout(videoPath string) (io.ReadCloser, error)
}

type YtDipVideoDownloader struct {
	Resolution string
}

func NewYtDipVideoDownloader(res string) *YtDipVideoDownloader {
	return &YtDipVideoDownloader{
		Resolution: res,
	}
}

func (y *YtDipVideoDownloader) DownloadToStdout(videoPath string) (io.ReadCloser, error) {
	ytDlpCmd := exec.Command(
		"yt-dlp",
		"-f", "worstvideo[height="+y.Resolution+"]",
		"--output", "-",
		videoPath,
	)

	stdout, err := ytDlpCmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	//	stderr, err := ytDlpCmd.StderrPipe()
	//	if err != nil {
	//		return nil, err
	//	}

	if err := ytDlpCmd.Start(); err != nil {
		return nil, err
	}

	//	go func() {
	//		scanner := bufio.NewScanner(stderr)
	//		for scanner.Scan() {
	//			log.Printf("yt-dlp Error %s", scanner.Text())
	//		}
	//	}()
	return stdout, nil
}
