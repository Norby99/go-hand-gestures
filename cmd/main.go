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
		win.IMShow(hr.DetectSkinColor(img))
		win.WaitKey(1)
	}
}
