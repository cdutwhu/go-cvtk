package main

import (
	"image"
	"image/color"
)

func FindPosByColor(img image.Image, c color.RGBA) (pos []image.Point) {
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
