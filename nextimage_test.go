package main

import (
	"fmt"
	"image"
	"sort"
	"testing"
)

func TestSearchNextROI(t *testing.T) {

	cfgEdge := "./cfg/AB-edge.json"
	inImage := "./in/sample/1.jpg"

	img := loadImg(inImage)
	centres := []image.Point{}

	for _, roi := range SearchNextROI(inImage, cfgEdge) {

		fmt.Println(roi.X, roi.Y)
		fmt.Println(roi.diff)

		gray := Cvt2Gray(roi.data)

		offset := gray.Stride * roiRadius
		line := gray.Pix[offset : offset+gray.Stride]

		ptr := []int{}
		for _, p := range line {
			ptr = append(ptr, int(p))
		}

		xMax, xUp, xDown := maxSlope(ptr, 7, 1)
		fmt.Println(xMax, xUp, xDown)

		centre := image.Point{X: roi.X - roiRadius + xUp, Y: roi.Y}
		centres = append(centres, centre)
	}

	sort.SliceStable(centres, func(i, j int) bool {
		return centres[i].Y < centres[j].Y
	})

	DrawCircle(img, centres, 2, "./out/next-dot.jpg")
	// DrawSpline(img, centres, 10, "./out/next-dot.jpg")
}
