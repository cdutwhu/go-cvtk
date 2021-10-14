package main

import (
	"fmt"
	"image"
	"image/color"

	gotkio "github.com/digisan/gotk/io"
)

func FindPosByClr(img image.Image, c color.RGBA) (pos []image.Point) {
	rect := img.Bounds()
	rgba := ROIrgba(img, rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
	for y := 0; y < rect.Dy(); y++ {
		start := y * rgba.Stride
		pln := rgba.Pix[start : start+rgba.Stride]
		for x := 0; x < rect.Dx(); x++ {
			p := pln[4*x:]
			cmp := color.RGBA{p[0], p[1], p[2], p[3]}
			if ColorEqual(c, cmp, 3, 0, 0, 256) {
				pos = append(pos, image.Point{x, y})
			}
		}
	}
	return
}

func FindROIrgbaByClr(img image.Image, c color.RGBA, sRadius int, auditPath string) (rgba []*image.RGBA) {
	gotkio.MustCreateDir(auditPath)
	for i, pos := range FindPosByClr(img, c) {
		roi := ROIrgbaV2(img, pos.X, pos.Y, sRadius)
		rgba = append(rgba, roi)
		if len(auditPath) > 0 {
			savePNG(roi, fmt.Sprintf("./%s/%00d-%d-%d.png", auditPath, i, pos.X, pos.Y))
		}
	}
	return
}
