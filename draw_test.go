package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"testing"
	"time"

	"github.com/digisan/gotk"
)

func TestDraw(t *testing.T) {
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

	imagerect := img.Bounds()
	fmt.Println(img.Bounds())

	left, top, right, bottom := imagerect.Min.X, imagerect.Min.Y, imagerect.Max.X, imagerect.Max.Y

	roi := ROIgray(img, left, top, right/2, bottom/2)

	roiClr := ROIrgba(img, left, top, right/2, bottom/2)
	savePNG(roiClr, "./out/color.png")

	roiCMYK := ROIcmyk(img, left, top, right/2, bottom/2)
	savePNG(roiCMYK, "./out/cmyk.png")

	roi1 := DrawHLine(roi, 30, 50, 500, nil)
	roi1 = DrawVLine(roi1, 40, 50, 1500, nil)
	roi1 = DrawRect(roi1, 60, 60, 300, 300, nil)

	savePNG(roi1, "./out/line.png")

	m, _, _ := histogram(roi1.Pix)
	hImg := DrawHisto(m, Peaks(m, 3, 1, -1), Bottoms(m, 3, 1, -1))
	savePNG(hImg, "./out/histo.png")
}
