package main

import (
	"testing"
)

func TestSearchNextROI(t *testing.T) {

	cfgEdge := "./cfg/AB-edge.json"

	inImage := "./in/sample/1.jpg"
	img := loadImg(inImage)
	keyPts := NextKeyPoints(inImage, cfgEdge, "")
	DrawCircle(img, keyPts, 3, "./out/next-dot-1.jpg")
	DrawSpline(img, keyPts, 10, "./out/next-dot-1.jpg")

	inImage = "./in/sample/2.jpg"
	img = loadImg(inImage)
	keyPts = NextKeyPoints(inImage, cfgEdge, "")
	DrawCircle(img, keyPts, 3, "./out/next-dot-2.jpg")
	DrawSpline(img, keyPts, 10, "./out/next-dot-2.jpg")

	inImage = "./in/sample/3.jpg"
	img = loadImg(inImage)
	keyPts = NextKeyPoints(inImage, cfgEdge, "")
	DrawCircle(img, keyPts, 3, "./out/next-dot-3.jpg")
	DrawSpline(img, keyPts, 10, "./out/next-dot-3.jpg")

	inImage = "./in/sample/4.jpg"
	img = loadImg(inImage)
	keyPts = NextKeyPoints(inImage, cfgEdge, "")
	DrawCircle(img, keyPts, 3, "./out/next-dot-4.jpg")
	DrawSpline(img, keyPts, 10, "./out/next-dot-4.jpg")
}
