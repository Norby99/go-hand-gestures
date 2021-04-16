package handrec

import (
	"image"
	"image/color"

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

// SkinMask
func SkinMask(img cv.Mat) cv.Mat {
	return cv.NewMat()
}

// SkinColorSampler draws two rectangles on the given image.
func SkinColorSampler(img cv.Mat) cv.Mat {
	height := img.Size()[1]
	width := img.Size()[0]
	rect1 := image.Rect(width/5, height/2, (width/5)+20, (height/2)+20)
	rect2 := image.Rect(width/5, height/3, (width/5)+20, (height/3)+20)
	cv.Rectangle(&img, rect1, color.RGBA{238, 11, 11, 1}, 3)
	cv.Rectangle(&img, rect2, color.RGBA{238, 11, 11, 1}, 3)
	return img
}
