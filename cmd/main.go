package main

import(
	"fmt"
	cv "gocv.io/x/gocv"
	"os"
)

func main() {

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
		key := window.WaitKey(1)
		if key == 27  || window.GetWindowProperty(cv.WindowPropertyAspectRatio) < 0{	// Close the window if ESC button was pressed or the x button is pressed
			os.Exit(3)
		}
	}

}
