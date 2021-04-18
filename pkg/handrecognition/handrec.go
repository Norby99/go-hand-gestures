package handrec

import (
	"image"
	"image/color"
	"log"
	"math"

	cv "gocv.io/x/gocv"
)

var faceClassifier = "../configs/haarcascade_frontalface_default.xml"

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

// RemoveBackground removes any unwanted elements in the given image.
func RemoveBackground(img cv.Mat) cv.Mat {
	// remove faces
	classifier := cv.NewCascadeClassifier()
	if !classifier.Load(faceClassifier) {
		log.Println("cannot load face classifier in path ", faceClassifier)
		return img
	}

	var faces []image.Rectangle
	frameGrey := cv.NewMat()
	cv.CvtColor(img, &frameGrey, cv.ColorBGRToGray)
	cv.EqualizeHist(frameGrey, &frameGrey)
	faces = classifier.DetectMultiScale(frameGrey)

	for _, face := range faces {
		cv.Rectangle(&img, face, color.RGBA{0, 0, 0, 0}, -1)
	}

	return img
}

// DetectSkinColor detects the user skin color and removes the rest.
func DetectSkinColor(img cv.Mat) cv.Mat {
	r1Height := img.Size()[1] / 2
	r2Height := img.Size()[1] / 3
	rWidth := img.Size()[0] / 4

	rect1 := image.Rect(rWidth, r1Height, rWidth+20, r1Height+20)
	rect2 := image.Rect(rWidth, r2Height, rWidth+20, r2Height+20)

	hueSat := cv.NewMat()
	cv.CvtColor(img, &hueSat, cv.ColorBGRToHSV)

	sample1 := hueSat.Region(rect1)
	sample2 := hueSat.Region(rect2)

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

	cv.Rectangle(&img, rect1, color.RGBA{238, 11, 11, 1}, 1)
	cv.Rectangle(&img, rect2, color.RGBA{238, 11, 11, 1}, 1)

	// handle noise
	morphStrElem := cv.GetStructuringElement(cv.MorphEllipse, image.Pt(3, 3))
	dilateStrElem := cv.GetStructuringElement(cv.MorphRect, image.Pt(3, 3))
	cv.MorphologyEx(img, &img, cv.MorphOpen, morphStrElem)
	cv.Dilate(img, &img, dilateStrElem)
	return img
}
