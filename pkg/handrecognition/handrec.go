package handrec

import (
	"image"
	"image/color"
	"math"

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
func DetectSkinColor(img cv.Mat) cv.Mat {
	r1Height := img.Size()[1] / 2
	r2Height := img.Size()[1] / 3
	rWidth := img.Size()[0] / 5

	rect1 := image.Rect(rWidth, r1Height, rWidth+20, r1Height+20)
	rect2 := image.Rect(rWidth, r2Height, rWidth+20, r2Height+20)
	cv.Rectangle(&img, rect1, color.RGBA{238, 11, 11, 1}, 1)
	cv.Rectangle(&img, rect2, color.RGBA{238, 11, 11, 1}, 3)

	hueSat := cv.NewMat()
	cv.CvtColor(img, &hueSat, cv.ColorBGRToHSV)

	sample1 := hueSat.Region(rect1)
	sample2 := hueSat.Region(rect2)

	// need to fix threshold values
	hueMeanSample1 := sample1.Mean()
	hueMeanSample2 := sample2.Mean()

	hLowThreshold := math.Min(hueMeanSample1.Val1, hueMeanSample2.Val1) - 80
	hHighThreshold := math.Max(hueMeanSample1.Val1, hueMeanSample2.Val1) + 30
	sLowThreshold := math.Min(hueMeanSample1.Val2, hueMeanSample2.Val2) - 80
	sHighThreshold := math.Max(hueMeanSample1.Val2, hueMeanSample2.Val2) + 30
	vLowThreshold := math.Min(hueMeanSample1.Val3, hueMeanSample2.Val3) - 80
	vHighThreshold := math.Max(hueMeanSample1.Val3, hueMeanSample2.Val3) + 30

	scalar1 := cv.NewScalar(hLowThreshold, sLowThreshold, vLowThreshold, 0)
	scalar2 := cv.NewScalar(hHighThreshold, sHighThreshold, vHighThreshold, 0)
	cv.InRangeWithScalar(hueSat, scalar1, scalar2, &img)
	return img
}
