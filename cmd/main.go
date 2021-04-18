package main

import (
	"fmt"

	hr "github.com/Norby99/go-hand-gestures/pkg/handrecognition"
	cv "gocv.io/x/gocv"
)

func main() {
	wc, err := cv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wc.Close()

	win := cv.NewWindow("go-hand-recognition")
	defer win.Close()
	img := cv.NewMat()
	defer img.Close()

	for {
		wc.Read(&img)
		img = hr.RemoveBackground(img)
		img = hr.DetectSkinColor(img)
		win.IMShow(img)
		win.WaitKey(1)
	}
}
