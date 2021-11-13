package main

import (
	"fmt"
	"image"
	"math"
	"sort"

	"github.com/digisan/gotk/slice/tu8i"
)

func peaksDiff(GRAY, RED, GREEN, BLUE []byte, gray, red, green, blue []byte) float64 {

	vDefault := 100.0

	sum, count := 0.0, 0
	for i, G := range GRAY {
		if len(gray) > i {
			g := gray[i]
			sum += math.Abs(float64(G) - float64(g))
			count++
		}
	}
	vGray := vDefault * wChPk[iGray]
	if count > 0 {
		vGray = sum / float64(count)
	}

	sum, count = 0.0, 0
	for i, R := range RED {
		if len(red) > i {
			r := red[i]
			sum += math.Abs(float64(R) - float64(r))
			count++
		}
	}
	vRed := vDefault * wChPk[iRed]
	if count > 0 {
		vRed = sum / float64(count)
	}

	sum, count = 0.0, 0
	for i, G := range GREEN {
		if len(green) > i {
			g := green[i]
			sum += math.Abs(float64(G) - float64(g))
			count++
		}
	}
	vGreen := vDefault * wChPk[iGreen]
	if count > 0 {
		vGreen = sum / float64(count)
	}

	sum, count = 0.0, 0
	for i, B := range BLUE {
		if len(blue) > i {
			b := blue[i]
			sum += math.Abs(float64(B) - float64(b))
			count++
		}
	}
	vBlue := vDefault * wChPk[iBlue]
	if count > 0 {
		vBlue = sum / float64(count)
	}

	return vGray*wChPk[iGray] + vRed*wChPk[iRed] + vGreen*wChPk[iGreen] + vBlue*wChPk[iBlue]
}

type ROICandidate struct {
	X, Y int
	diff float64
	data *image.RGBA
}

func searchNextROI(imgPath, cfgEdge string) (selected []ROICandidate) {

	img := loadImg(imgPath)
	edge := LoadLastRecord(cfgEdge)

	for _, pt := range edge.Pts {
		// out
		candidates := []ROICandidate{}

		start := pt.X - searchOffset
		end := pt.X + searchOffset
		for s := start; s < end; s++ {

			roi := ROIrgbaV2(img, s, pt.Y, roiRadius)
			r, g, b, _ := SplitRGBA(roi)
			gray := Cvt2Gray(roi)

			// out
			chPeaks := [][]byte{}

			for _, ch := range []*image.Gray{gray, r, g, b} {

				m, _, _ := histogram(ch.Pix)                                             // histogram data
				peaks := Peaks(m, 3, 1, 2)                                               // only find 2 peaks
				ks, _ := tu8i.Map2KVs(peaks, func(i, j byte) bool { return i > j }, nil) // DESC, most background value at the first

				chPeaks = append(chPeaks, ks)
			}

			candidates = append(candidates, ROICandidate{
				X: s,
				Y: pt.Y,
				diff: peaksDiff(
					pt.GrayPeaks, pt.RedPeaks, pt.GreenPeaks, pt.BluePeaks,
					chPeaks[iGray], chPeaks[iRed], chPeaks[iGreen], chPeaks[iBlue],
				),
				data: roi,
			})
		}

		sort.SliceStable(candidates, func(i, j int) bool {
			return candidates[i].diff < candidates[j].diff
		})

		selected = append(selected, candidates[0])
	}

	return
}

func makeNextEdgeCfg(selected []ROICandidate, cfgEdge, recordName, imgPath string) {
	record := NewEdgeRecord(recordName, imgPath)
	for _, roi := range selected {
		r, g, b, _ := SplitRGBA(roi.data)
		gray := Cvt2Gray(roi.data)
		chPeaks := [][]byte{}
		for _, ch := range []*image.Gray{gray, r, g, b} {
			m, _, _ := histogram(ch.Pix)                                             // histogram data
			peaks := Peaks(m, 3, 1, 2)                                               // only find 2 peaks
			ks, _ := tu8i.Map2KVs(peaks, func(i, j byte) bool { return i > j }, nil) // desc, most background value at the first
			chPeaks = append(chPeaks, ks)
		}
		record.AddPtInfo(roi.X, roi.Y, chPeaks[iGray], chPeaks[iRed], chPeaks[iGreen], chPeaks[iBlue])
	}
	record.Log(cfgEdge)
}

func NextKeyPoints(imgPath, cfgEdge, nextRecordName string) (centres []image.Point) {

	selected := searchNextROI(imgPath, cfgEdge)

	for _, roi := range selected {

		fmt.Println(roi.X, roi.Y)
		fmt.Println(roi.diff)

		gray := Cvt2Gray(roi.data)

		ptr := GrayStripeV(gray, roiRadius)

		// offset := gray.Stride * roiRadius
		// line := gray.Pix[offset : offset+gray.Stride]
		// ptr := []int{}
		// for _, p := range line {
		// 	ptr = append(ptr, int(p))
		// }

		yMax, yUp, yDown := maxSlope(ptr, 7, 1)
		fmt.Println(yMax, yUp, yDown)

		// centre := image.Point{X: roi.X - roiRadius + xUp, Y: roi.Y}
		centre := image.Point{
			X: roi.X,
			Y: roi.Y - roiRadius + yDown,
		}
		centres = append(centres, centre)
	}

	sort.SliceStable(centres, func(i, j int) bool {
		return centres[i].Y < centres[j].Y
	})

	// [selected] could be modified
	makeNextEdgeCfg(selected, cfgEdge, nextRecordName, imgPath)

	return
}
