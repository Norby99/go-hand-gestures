package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	cv "gocv.io/x/gocv"
	//py "../pkg/pythonGoInterface"
	"os/exec"
)

/*
?I'll leave this here, so in case one day you decide to continue it, I'll have some info.
?1) I can't import the file from pythonGoInterface
?2) I can't execute the python code, cuz the string is too long, maybe writing the string to a text file will be the solution
*/

func toString(mat cv.Mat) (string) {
	img, _ := mat.ToImage()
	bounds := img.Bounds()
	//stringMatrix := ""
	var stringMatrix strings.Builder

	//don't get surprised of reversed order everywhere below
	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			stringMatrix.WriteString(strconv.FormatUint(uint64(byte(r>>8)), 10))
			stringMatrix.WriteString(" ")
			stringMatrix.WriteString(strconv.FormatUint(uint64(byte(g>>8)), 10))
			stringMatrix.WriteString(" ")
			stringMatrix.WriteString(strconv.FormatUint(uint64(byte(b>>8)), 10))
		}
	}
	return stringMatrix.String()
}

func main() {

	pythonPath := "import os, sys;currentdir = os.path.dirname(os.path.realpath(os.getcwd()));parentdir = os.path.dirname(currentdir);sys.path.append(parentdir);from pkg.pythonGoInterface import pythonGoInterface"	// I don't really know how i wrote this

	window := cv.NewWindow("Window")
	img := cv.IMRead("campione.png", cv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", "campione.png")
		return
	}

	myOutput := toString(img)

	f, err := os.Create("data.txt")

    if err != nil {
        fmt.Print(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(myOutput)

    if err2 != nil {
        fmt.Print(err2)
    }

	//TODO this code is perfect, need to be finished
	cmd := exec.Command("python.exe",  "-c", pythonPath+"; print(pythonGoInterface.draw_styled_landmarks("+myOutput+"))")
	out, err := cmd.CombinedOutput()
	if err != nil { fmt.Println(err); }
	fmt.Println(string(out))

	for{
		window.IMShow(img)
		key := window.WaitKey(1)
		if key == 27  || window.GetWindowProperty(cv.WindowPropertyAspectRatio) < 0{	// Close the window if ESC button was pressed or the x button is pressed
			os.Exit(3)
		}
	}

/*
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
*/

		/*
		TODO this code is perfect, need to be finished
		cmd := exec.Command("python.exe",  "-c", pythonPath+"; print(pythonGoInterface.myFunction("+"))")
		out, err := cmd.CombinedOutput()
    	if err != nil { fmt.Println(err); }
    	fmt.Println(string(out))*/
/*
		os.Exit(3)

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		key := window.WaitKey(1)
		if key == 27  || window.GetWindowProperty(cv.WindowPropertyAspectRatio) < 0{	// Close the window if ESC button was pressed or the x button is pressed
			os.Exit(3)
		}
	}*/

}
