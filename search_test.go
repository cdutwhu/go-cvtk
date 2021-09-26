package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"testing"
	"time"

	"github.com/digisan/gotk"
)

func TestFindPosByColor(t *testing.T) {
	defer gotk.TrackTime(time.Now())

	f, err := os.Open("./in/calibrate.JPG")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fmtName)

	pos := FindPosByColor(img, color.RGBA{254, 0, 0, 255})
	fmt.Println(len(pos), pos)
	for _, p := range pos {
		roi := ROIrgbaV2(img, p.X, p.Y, 10)
		savePNG(roi, "./out/roi.png")
	}
}
