package handrec

import (
	"image"

	cv "gocv.io/x/gocv"
)

// HandPos contains the hand's information in the current frame.
type HandPos struct {
	// TODO
	/*
		index  image.Point
		middle image.Point
		ring   image.Point
		little image.Point
	*/
}

// Detect searchs for a hand in the given image, if nothing is found an error is returned (?)
func Detect(cap cv.Mat) (HandPos, error) {
	return HandPos{}, nil
}

/*
	per ora sto cercando di tradurre sto bocchino fatto in python in go
	https://github.com/madhav727/medium/blob/master/finger_counting_video.py
	se trovate qualcosa di meglio ditelo non fate gli infami
*/

// SkinMask
/* note:
Skin mask dovrebbe "rimuovere" la pelle o texture inutili lasciando una imamgine
in bianco e nero con la mano più facilmente visibile.
*/
func SkinMask(img cv.Mat) cv.Mat {
	hueSat := cv.NewMat()
	cv.CvtColor(img, &hueSat, cv.ColorRGBToHSV)
	skinRegionHS := cv.NewMat()
	// TODO
	/*
		issue with RGB values to be converted to BGR values.
		lower and upper should also be made constants (?)
	*/
	lower, _ := cv.NewMatFromBytes(1, 3, cv.MatTypeCV8U, []byte{0, 48, 80})
	upper, _ := cv.NewMatFromBytes(1, 3, cv.MatTypeCV8U, []byte{20, 255, 255})
	cv.InRange(hueSat, lower, upper, &skinRegionHS)
	blur := cv.NewMat()
	cv.Blur(skinRegionHS, &blur, image.Point{2, 2})
	cv.Threshold(blur, &img, 0, 255, cv.ThresholdBinary)
	return img
}
