package main

func main() {

	//? This code detects a face :D
	/*
		// parse args
		deviceID := 0
		xmlFile := "../data/haarcascade_frontalface_default.xml"

		// open webcam
		webcam, err := gocv.OpenVideoCapture(deviceID)
		if err != nil {
			fmt.Printf("error opening video capture device: %v\n", deviceID)
			return
		}
		defer webcam.Close()

		// open display window
		window := gocv.NewWindow("Face Detect")
		defer window.Close()

		// prepare image matrix
		img := gocv.NewMat()
		defer img.Close()

		// color for the rect when faces detected
		blue := color.RGBA{0, 0, 255, 0}

		// load classifier to recognize faces
		classifier := gocv.NewCascadeClassifier()
		defer classifier.Close()

		if !classifier.Load(xmlFile) {
			fmt.Printf("Error reading cascade file: %v\n", xmlFile)
			return
		}

		fmt.Printf("Start reading device: %v\n", deviceID)
		for {
			if ok := webcam.Read(&img); !ok {
				fmt.Printf("Device closed: %v\n", deviceID)
				return
			}
			if img.Empty() {
				continue
			}

			// detect faces
			rects := classifier.DetectMultiScale(img)
			fmt.Printf("found %d faces\n", len(rects))

			// draw a rectangle around each face on the original image,
			// along with text identifing as "Human"
			for _, r := range rects {
				gocv.Rectangle(&img, r, blue, 3)

				size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
				pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
				gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
			}

			// show the image in the window, and wait 1 millisecond
			window.IMShow(img)
			if window.WaitKey(1) >= 0 {
				break
			}
		}*/

	//? This code loads an image and displays it on a new window
	/*window := gocv.NewWindow("Hello")
	img := gocv.IMRead("../hand-1.png", gocv.IMReadColor)

	for {
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}*/

}
