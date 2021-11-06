package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"testing"
	"time"

	"github.com/digisan/gotk"
)

func TestSplitRGBA(t *testing.T) {
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
