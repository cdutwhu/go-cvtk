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
	"github.com/digisan/gotk/slice/tu8i"
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

	mChPeak := make(map[string][][]byte) // e.g. "R" [[87, 120], [87, 114] ... ]

	mPeak0Vals := make(map[string][]byte)
	mPeak1Vals := make(map[string][]byte)

	mPtRGBA := FindROIrgbaByClr(img, color.RGBA{254, 0, 0, 255}, 20, "./out/")

	for pt, roi := range mPtRGBA {

		r, g, b, _ := SplitRGBA(roi)
		gray := Cvt2Gray(roi)

		desChClr := []string{"Gray", "R", "G", "B"}
		for iCh, ch := range []*image.Gray{gray, r, g, b} {

			m, _, _ := histogram(ch.Pix)

			peaks := Peaks(m, 3, 1, 2)
			// fmt.Println("peak:", peaks)
			ks, _ := tu8i.Map2KVs(peaks, func(i, j byte) bool { return i < j }, nil)

			clr := desChClr[iCh]
			mChPeak[clr] = append(mChPeak[clr], ks)

			fmt.Println(pt, clr, ks)

			mPeak0Vals[clr] = append(mPeak0Vals[clr], ks[0])
			mPeak1Vals[clr] = append(mPeak1Vals[clr], ks[1])

			// bottoms := Bottoms(m, 3, 1, 1)
			// fmt.Println("bottom:", bottoms)

			hImg := DrawHisto(m, peaks, nil)
			savePNG(hImg, fmt.Sprintf("./out/histo-%v-%s.png", pt, clr))
		}
		fmt.Println()
	}

	fmt.Println("------------------------------------")
	for k, v := range mPeak0Vals {
		fmt.Println(k, v)
	}
	fmt.Println("------------------------------------")
	for k, v := range mPeak1Vals {
		fmt.Println(k, v)
	}

}
