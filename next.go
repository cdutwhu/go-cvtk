package main

import (
	"image"
	"math"
	"sort"
)

type ROICandidate struct {
	pt   image.Point
	data *image.RGBA
	dy   int
}

func nextROICandidates(imgPath, cfgEdge string) (selected []ROICandidate) {
	img := loadImg(imgPath)
	edge := LoadLastRecord(cfgEdge)
	for _, pt := range edge.Points() {
		selected = append(selected, ROICandidate{
			pt:   pt,
			data: ROIrgbaV2(img, pt.X, pt.Y, roiRadius),
		})
	}
	sort.SliceStable(selected, func(i, j int) bool {
		return selected[i].pt.X < selected[j].pt.X
	})
	return
}

func makeNextEdgeCfg(selected []ROICandidate, cfgEdge, recordName, imgPath string) {
	record := NewEdgeRecord(recordName, imgPath)
	for _, roi := range selected {
		f := feature(roi.data)
		above, below, left, right := f[0], f[1], f[2], f[3]
		record.AddPtInfo(roi.pt.X, roi.pt.Y, above, below, left, right)
	}
	record.Log(cfgEdge)
}

func NextKeyPoints(imgPath, cfgEdge, nextRecordName string) (centres []image.Point) {

	img := loadImg(imgPath)
	edge := LoadLastRecord(cfgEdge)
	selected := nextROICandidates(imgPath, cfgEdge)

	pts4all := []ROICandidate{}

	for _, roi := range selected {

		pts4each := []ROICandidate{}

		// looking for edge config
		for _, pt := range edge.Pts {

			// refer to suitable config roi
			if roi.pt.X == pt.X {

				// gray := Cvt2Gray(roi.data)
				r, _, _, _ := SplitRGBA(roi.data) // choose [red] channel for slope
				ptr := GrayStripeV(r, roiRadius)
				ps := slope(ptr, slopeStep, 0)

				// if pt.ValAbove < pt.ValBelow {
				// 	// up -> down : dark -> bright
				// 	for _, s := range ps[:5] {

				// 		y := roi.Y - roiRadius + s
				// 		roi := ROIgrayV2(img, pt.X, y, roiRadius)
				// 		f := feature(roi)
				// 		above, below := f[0], f[1]

				// 		if math.Abs(float64(pt.ValAbove)-float64(above)) < ERR &&
				// 			math.Abs(float64(pt.ValBelow)-float64(below)) < ERR {
				// 			centre := image.Point{
				// 				X: pt.X,
				// 				Y: y,
				// 			}
				// 			centres = append(centres, centre)
				// 			// dYs = append(dYs, math.Abs(float64(centre.Y)-float64(pt.Y)))

				// 			continue NEXT
				// 		}
				// 	}
				// }

				if pt.ValAbove > pt.ValBelow {
					// up -> down : bright -> dark
					for i := len(ps) - 1; i >= len(ps)-5; i-- {
						s := ps[i]

						y := roi.pt.Y - roiRadius + s
						tempROI := ROIrgbaV2(img, pt.X, y, roiRadius)
						f := feature(tempROI)
						above, below := f[0], f[1]

						if math.Abs(float64(pt.ValAbove)-float64(above)) < ERR &&
							math.Abs(float64(pt.ValBelow)-float64(below)) < ERR {
							pts4each = append(pts4each, ROICandidate{
								pt: image.Point{
									X: pt.X,
									Y: y,
								},
								data: tempROI,
								dy:   int(math.Abs(float64(y) - float64(pt.Y))),
							})
						}
					}
				}
			}
		}

		sort.SliceStable(pts4each, func(i, j int) bool {
			return pts4each[i].dy < pts4each[j].dy
		})

		if len(pts4each) > 0 {
			wanted := pts4each[0]
			centres = append(centres, wanted.pt)
			pts4all = append(pts4all, ROICandidate{
				pt:   wanted.pt,
				data: wanted.data,
			})
		}
	}

	// [pts4all] for next config
	makeNextEdgeCfg(pts4all, cfgEdge, nextRecordName, imgPath)
	return
}
