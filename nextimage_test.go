package main

import (
	"fmt"
	"testing"
)

func TestSearchNextROI(t *testing.T) {

	cfgEdge := "./cfg/AB-edge.json"
	inImage := "./in/sample/1.jpg"

	rois := SearchNextROI(inImage, cfgEdge)

	for _, roi := range rois {
		fmt.Println(roi.X, roi.Y)
		fmt.Println(roi.diff)
	}
}
