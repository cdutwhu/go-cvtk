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

	f, err := os.Open("./in/mel.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fmtName)

	mPtRGBA := FindROIrgbaByClr(img, color.RGBA{254, 0, 0, 255}, 20, "./out/")
	fmt.Println(mPtRGBA)
}
