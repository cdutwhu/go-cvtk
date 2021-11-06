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

	f, err := os.Open("./in/start/1.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fmtName)

	mPtROI := FindROIrgbaByClr(img, color.RGBA{254, 0, 0, 255}, 70, 8, "./out/")
	fmt.Println(mPtROI)
}

// func TestFindROIrgbaByBlob(t *testing.T) {
// 	defer gotk.TrackTime(time.Now())

// 	f, err := os.Open("./in/start/1.jpg")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer f.Close()

// 	img, fmtName, err := image.Decode(f)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(fmtName)

// 	mPtRGBA := FindROIrgbaByBlob(img, 70,
// 		func(x, y int, p byte) bool {
// 			return p > 253
// 		}, func(x, y int, p byte) bool {
// 			return p < 2
// 		}, func(x, y int, p byte) bool {
// 			return p < 2
// 		}, 5, "./out/")
// 	fmt.Println(mPtRGBA)
// }
