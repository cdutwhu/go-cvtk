package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/digisan/gotk/slice/tu8i"
)

const (
	iGray = iota
	iRed
	iGreen
	iBlue
)

var (
	chClr = []string{"Gray", "Red", "Green", "Blue"}
)

func BuildModel(recordPath, recordName, imagePath string, aim color.RGBA) {

	record := NewEdgeRecord(recordName, imagePath)

	f, err := os.Open(imagePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fmtName)

	mPtROI := FindROIrgbaByClr(img, aim, 70, 7, "./out/")
	fmt.Println(mPtROI)

	for pt, roi := range mPtROI {

		r, g, b, _ := SplitRGBA(roi)
		gray := Cvt2Gray(roi)

		chPeaks := [][]byte{}

		for iCh, ch := range []*image.Gray{gray, r, g, b} {

			clr := chClr[iCh]            // color name
			m, _, _ := histogram(ch.Pix) // histogram data

			peaks := Peaks(m, 3, 1, 2) // only find 2 peaks
			// fmt.Println("peak:", peaks)
			// savePNG(DrawHisto(m, peaks, nil), fmt.Sprintf("./out/histo-%v-%s.png", pt, clr)) // audit

			ks, _ := tu8i.Map2KVs(peaks, func(i, j byte) bool { return i > j }, nil) // desc, most background value at the first
			fmt.Println(pt, clr, "----- two peaks pos:", ks)

			chPeaks = append(chPeaks, ks)
		}

		record.AddPtInfo(pt.X, pt.Y, chPeaks[iGray], chPeaks[iRed], chPeaks[iGreen], chPeaks[iBlue])

		fmt.Println()
	}

	record.Log(recordPath)
}
