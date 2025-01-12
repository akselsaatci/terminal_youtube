package renderer

import (
	"fmt"
	"math"
	"time"
)

// TODO not sure if i need it or not
type Renderer interface {
	Render()
}

type TerminalRenderer struct {
	framePerSecond    int
	timeBetweenFrames float64
	framesChannel     <-chan string
	doneChannel       chan<- bool
	frameDoneChannel  <-chan bool
}

func NewTerminalRenderer(fps int, framesChannel <-chan string, doneChannel chan<- bool, frameDoneChannel <-chan bool) *TerminalRenderer {
	return &TerminalRenderer{
		framePerSecond: fps,
		framesChannel:  framesChannel,
		// TODO maybe can convert it to ms
		timeBetweenFrames: float64(1) / float64(fps),
		frameDoneChannel:  frameDoneChannel,
		doneChannel:       doneChannel,
	}

}
func (t *TerminalRenderer) Render() {
	//this is for calculating the sleep time between frames
	startTime := time.Now()
	//these are for tracking fps
	fpsStartTime := time.Now()
	frameCount := 0

	isDone := false
	go func() {
		isDone = <-t.frameDoneChannel
	}()

	//main for loop
	for {
		//clear code
		fmt.Print("\033[H\033[2J")
		currentFrame := <-t.framesChannel
		if currentFrame == "%END%" && isDone {
			t.doneChannel <- true
			close(t.doneChannel)
			return
		}

		fmt.Println(currentFrame)
		frameCount++
		fpsElapsed := time.Since(fpsStartTime).Seconds()
		//to calculate fps in every frame
		if fpsElapsed >= 1.0 {
			currentFPS := float64(frameCount) / fpsElapsed
			fmt.Printf("\n\t\t\t\t\t\nFPS: %.2f\n", currentFPS)
			fpsStartTime = time.Now()
			frameCount = 0
		}
		elapsed := time.Since(startTime).Seconds()
		// if sleep needed we decide it here
		sleepTime := math.Max(0, t.timeBetweenFrames-elapsed)
		time.Sleep(time.Duration(sleepTime * float64(time.Second)))
		startTime = time.Now()
	}

}
