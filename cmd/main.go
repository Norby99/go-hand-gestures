package main

import(
	"fmt"
	cv "gocv.io/x/gocv"
	"os"
	//"os/exec"
)

func toString(mat cv.Mat) (string) {
	img, _ := mat.ToImage()
	bounds := img.Bounds()
	stringMatrix := ""

	//don't get surprised of reversed order everywhere below
	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			stringMatrix += string(byte(b>>8)) + string(byte(g>>8)) + string(byte(r>>8))	// ! absolutly too slow
		}
	}
	return stringMatrix
}

func main() {

	//pythonPath := "import os, sys;currentdir = os.path.dirname(os.path.realpath(os.getcwd()));parentdir = os.path.dirname(currentdir);sys.path.append(parentdir);from pkg.pythonGoInterface import pythonfile"	// I don't really know how i wrote this

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

		myString := toString(img)
		fmt.Print("ciao")
		fmt.Printf("\n-%s\n\n", myString[0:9])


		/*
		TODO this code is perfect, need to be finished
		cmd := exec.Command("python.exe",  "-c", pythonPath+"; print(pythonGoInterface.myFunction("+"))")
		out, err := cmd.CombinedOutput()
    	if err != nil { fmt.Println(err); }
    	fmt.Println(string(out))*/

		os.Exit(3)

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		key := window.WaitKey(1)
		if key == 27  || window.GetWindowProperty(cv.WindowPropertyAspectRatio) < 0{	// Close the window if ESC button was pressed or the x button is pressed
			os.Exit(3)
		}
	}

}
