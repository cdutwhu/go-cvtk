package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/digisan/gotk/slice/tu8i"
)

func BuildModel(imagepath, recordpath string, aim color.RGBA) {

	f, err := os.Open(imagepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	record := NewEdgeRecord("AB", imagepath)

	img, fmtName, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fmtName)

	mPtROI := FindROIrgbaByClr(img, aim, 70, 7, "./out/")
	fmt.Println(mPtROI)

	mChPeak := make(map[string][][]byte) // e.g. "R" [[87, 120], [87, 114] ... ] for list of 2 peaks
	// mPeak0Vals := make(map[string][]byte)
	// mPeak1Vals := make(map[string][]byte)

	desChClr := []string{"Gray", "R", "G", "B"}

	for pt, roi := range mPtROI {

		r, g, b, _ := SplitRGBA(roi)
		gray := Cvt2Gray(roi)

		chPeaks := [][]byte{}

		for iCh, ch := range []*image.Gray{gray, r, g, b} {

			m, _, _ := histogram(ch.Pix)

			peaks := Peaks(m, 3, 1, 2) // only find 2 peaks
			// fmt.Println("peak:", peaks)

			// bottoms := Bottoms(m, 3, 1, 1)
			// fmt.Println("bottom:", bottoms)

			ks, _ := tu8i.Map2KVs(peaks, func(i, j byte) bool { return i > j }, nil) // desc, most background value at the first

			clr := desChClr[iCh]
			mChPeak[clr] = append(mChPeak[clr], ks)

			fmt.Println(pt, clr, "----- two peaks pos:", ks)

			chPeaks = append(chPeaks, ks)

			hImg := DrawHisto(m, peaks, nil)
			savePNG(hImg, fmt.Sprintf("./out/histo-%v-%s.png", pt, clr))

			// if len(ks) > 0 {
			// 	mPeak0Vals[clr] = append(mPeak0Vals[clr], ks[0])
			// 	if len(ks) > 1 {
			// 		mPeak1Vals[clr] = append(mPeak1Vals[clr], ks[1])
			// 	}
			// }
		}

		record.AddPtInfo(pt.X, pt.Y, chPeaks[0], chPeaks[1], chPeaks[2], chPeaks[3])

		fmt.Println()
	}

	record.Log(recordpath)

	// fmt.Println("Peak 1 pos for each ROI:")
	// fmt.Println("------------------------------------")
	// for k, v := range mPeak0Vals {
	// 	fmt.Println(k, v)
	// }

	// fmt.Println("")

	// fmt.Println("Peak 2 pos for each ROI:")
	// fmt.Println("------------------------------------")
	// for k, v := range mPeak1Vals {
	// 	fmt.Println(k, v)
	// }
}
