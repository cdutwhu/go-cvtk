package main

import (
	"image"
	"image/draw"
)

func ROIrgba(img image.Image, left, top, right, bottom int) *image.RGBA {
	rect := image.Rect(left, top, right, bottom)
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, image.Point{left, top}, draw.Src)
	return rgba
}

func ROIrgbaV2(img image.Image, cx, cy, sRadius int) *image.RGBA {
	left := cx - sRadius
	top := cy - sRadius
	right := cx + sRadius
	bottom := cy + sRadius
	return ROIrgba(img, left, top, right, bottom)
}

func ROIcmyk(img image.Image, left, top, right, bottom int) *image.CMYK {
	rect := image.Rect(left, top, right, bottom)
	cmyb := image.NewCMYK(rect)
	draw.Draw(cmyb, rect, img, image.Point{left, top}, draw.Src)
	return cmyb
}

func ROIcmykV2(img image.Image, cx, cy, sRadius int) *image.CMYK {
	left := cx - sRadius
	top := cy - sRadius
	right := cx + sRadius
	bottom := cy + sRadius
	return ROIcmyk(img, left, top, right, bottom)
}

func ROIgray(img image.Image, left, top, right, bottom int) *image.Gray {
	rect := image.Rect(left, top, right, bottom)
	gray := image.NewGray(rect)
	draw.Draw(gray, rect, img, image.Point{left, top}, draw.Src)
	return gray
}

func ROIgrayV2(img image.Image, cx, cy, sRadius int) *image.Gray {
	left := cx - sRadius
	top := cy - sRadius
	right := cx + sRadius
	bottom := cy + sRadius
	return ROIgray(img, left, top, right, bottom)
}

func Cvt2Gray(img image.Image) *image.Gray {
	rect := img.Bounds()
	gray := image.NewGray(rect)
	draw.Draw(gray, rect, img, rect.Min, draw.Src)
	return gray
}
