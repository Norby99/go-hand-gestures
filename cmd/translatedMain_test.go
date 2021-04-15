package main

import (
	"fmt"
	"testing"
	cv "gocv.io/x/gocv"
)

func TestMain(t *testing.T) {

	// open webcam
	webcam, err := cv.OpenVideoCapture(0)
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", 0)
		return
	}
	defer webcam.Close()

	// open display window
	window := cv.NewWindow("Window")

	img := cv.NewMat()
		defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed\n")
			return
		}
		if img.Empty() {
			continue
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

}