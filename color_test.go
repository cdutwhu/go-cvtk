package main

import (
	_ "image/png"
	"testing"
	"time"

	"github.com/digisan/gotk"
)

func TestSplitRGBA(t *testing.T) {
	defer gotk.TrackTime(time.Now())

	img := loadImg("./in/mel.png")

	// SplitRGBA(img)

	r, g, b, a := SplitRGBA(img)
	savePNG(r, "./out/r.png")
	savePNG(g, "./out/g.png")
	savePNG(b, "./out/b.png")
	savePNG(a, "./out/a.png")

	///

	com := CompositeRGBA(r, g, b, a)
	savePNG(com, "./out/MelCom.png")
}
