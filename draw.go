package main

import (
	"image"
	"image/draw"

	"github.com/digisan/go-handy-cv/blob"
	"github.com/digisan/gotk/slice/ti"
	"github.com/digisan/gotk/slice/tu8i"
)

func DrawRect(img *image.Gray, left, top, right, bottom int, paint *image.Gray) *image.Gray {
	if paint == nil {
		paint = image.NewGray(image.Rect(0, 0, right, bottom))
		for i := 0; i < len(paint.Pix); i++ {
			paint.Pix[i] = 0
		}
	}

	draw.Draw(img, image.Rect(left, top, right, top+1), paint, image.Point{0, 0}, draw.Src)
	draw.Draw(img, image.Rect(left, top, left+1, bottom), paint, image.Point{0, 0}, draw.Src)
	draw.Draw(img, image.Rect(left, bottom, right, bottom+1), paint, image.Point{0, 0}, draw.Src)
	draw.Draw(img, image.Rect(right, top, right+1, bottom+1), paint, image.Point{0, 0}, draw.Src)
	return img
}

func DrawHLine(img *image.Gray, y, left, right int, paint *image.Gray) *image.Gray {
	if paint == nil {
		paint = image.NewGray(image.Rect(0, 0, right, 1))
		for i := 0; i < len(paint.Pix); i++ {
			paint.Pix[i] = 0
		}
	}

	draw.Draw(img, image.Rect(left, y, right, y+1), paint, image.Point{0, 0}, draw.Src)
	return img
}

func DrawVLine(img *image.Gray, x, top, bottom int, paint *image.Gray) *image.Gray {
	if paint == nil {
		paint = image.NewGray(image.Rect(0, 0, 1, bottom))
		for i := 0; i < len(paint.Pix); i++ {
			paint.Pix[i] = 0
		}
	}

	draw.Draw(img, image.Rect(x, top, x+1, bottom), paint, image.Point{0, 0}, draw.Src)
	return img
}

func DrawBlob(left, top, right, bottom int, bytes []byte) *image.Gray {
	paint := image.NewGray(image.Rect(left, top, right, bottom))
	for i := 0; i < len(paint.Pix); i++ {
		paint.Pix[i] = 0
	}

	blobs := blob.DetectBlob(right-left, bottom-top, right-left, bytes, func(x, y int, p byte) bool {
		return p < 40
	})

	hImg := image.NewGray(image.Rect(left, top, right, bottom))
	hImg.Pix = bytes

	for _, blob := range blobs {
		loc := blob.Loc()
		left, top, right, bottom := loc[0].X, loc[0].Y, loc[1].X, loc[1].Y
		hImg = DrawRect(hImg, left, top, right, bottom, paint)
	}
	return hImg
}

func DrawHisto(mHisto, mPeak, mBottom map[byte]int) (hImg *image.Gray) {

	_, vs := tu8i.Map2KVs(mHisto, nil, nil)
	maxCnt := ti.Max(vs...)
	r := float64(maxCnt) / float64(255)
	hImg = image.NewGray(image.Rect(0, 0, 256, 256))

	// drawing
	mY := make(map[byte]int)
	for k, v := range mHisto {
		mY[k] = int(float64(v) / r)
	}
	ks, vs := tu8i.Map2KVs(mY, func(i, j byte) bool { return i < j }, nil)
	vs = smooth(vs) // remove noise

	paint := image.NewGray(image.Rect(0, 0, 1, 256))
	for i := 0; i < len(paint.Pix); i++ {
		paint.Pix[i] = 255
	}
	for i := 0; i < len(ks); i++ {
		k, v := ks[i], vs[i]
		DrawVLine(hImg, int(k), 256-v, 256, paint)
	}

	// mark peak
	if len(mPeak) > 0 {
		paintPeak := image.NewGray(image.Rect(0, 0, 1, 50))
		for i := 0; i < len(paintPeak.Pix); i++ {
			paintPeak.Pix[i] = 50
		}
		for x, y := range mPeak {
			y = int(float64(y) / r)
			DrawVLine(hImg, int(x), 256-y+5, 256-y+30, paintPeak)
		}
	}

	// mark bottom
	if len(mBottom) > 0 {
		paintBottom := image.NewGray(image.Rect(0, 0, 1, 50))
		for i := 0; i < len(paintBottom.Pix); i++ {
			paintBottom.Pix[i] = 200
		}
		for x, y := range mBottom {
			y = int(float64(y) / r)
			DrawVLine(hImg, int(x), 256-y-30, 256-y-5, paintBottom)
		}
	}

	return hImg
}
