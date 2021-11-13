package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"testing"
)

func TestSearchNextROI(t *testing.T) {

	os.Mkdir("./cfg/", os.ModePerm)

	modelImage := "./in/sample/std.jpg"

	cfgEdgeAB := "./cfg/AB-edge.json"
	cfgEdgeBC := "./cfg/BC-edge.json"

	BuildModel(cfgEdgeAB, "AB", modelImage, color.RGBA{255, 0, 0, 255})
	BuildModel(cfgEdgeBC, "BC", modelImage, color.RGBA{0, 255, 0, 255})

	///////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////

	mode := "DOT"

	var draw func(img image.Image, centres []image.Point, r int, color string, savePath string) image.Image

	switch mode {
	case "DOT":
		draw = DrawCircle
	case "LINE":
		draw = DrawSpline
	default:
		draw = DrawSpline
	}

	inImage := "./in/sample/1.jpg"
	outImage := "./out/next-dot-1.jpg"

	cfgEdge := "./cfg/AB-edge.json"
	color := "R"
	rs := 5 // *** [dot-radius] or [line-step] ***

	img := loadImg(inImage)
	keyPts := NextKeyPoints(inImage, cfgEdge, "")
	draw(img, keyPts, rs, color, outImage)

	cfgEdge = "./cfg/BC-edge.json"
	color = "G"

	inImage = outImage
	img = loadImg(inImage)
	keyPts = NextKeyPoints(inImage, cfgEdge, "")
	draw(img, keyPts, rs, color, inImage)

	fmt.Println()

	///////////////////////////////////////////////////

	inImage = "./in/sample/2.jpg"
	outImage = "./out/next-dot-2.jpg"

	cfgEdge = "./cfg/AB-edge.json"
	color = "R"

	img = loadImg(inImage)
	keyPts = NextKeyPoints(inImage, cfgEdge, "")
	draw(img, keyPts, rs, color, outImage)

	cfgEdge = "./cfg/BC-edge.json"
	color = "G"

	inImage = outImage
	img = loadImg(inImage)
	keyPts = NextKeyPoints(inImage, cfgEdge, "")
	draw(img, keyPts, rs, color, inImage)

	fmt.Println()

	///////////////////////////////////////////////////

	inImage = "./in/sample/3.jpg"
	outImage = "./out/next-dot-3.jpg"

	cfgEdge = "./cfg/AB-edge.json"
	color = "R"

	img = loadImg(inImage)
	keyPts = NextKeyPoints(inImage, cfgEdge, "")
	draw(img, keyPts, rs, color, outImage)

	cfgEdge = "./cfg/BC-edge.json"
	color = "G"

	inImage = outImage
	img = loadImg(inImage)
	keyPts = NextKeyPoints(inImage, cfgEdge, "")
	draw(img, keyPts, rs, color, inImage)

	fmt.Println()

	//////////////////////////////////////////////////////////

	inImage = "./in/sample/4.jpg"
	outImage = "./out/next-dot-4.jpg"

	cfgEdge = "./cfg/AB-edge.json"
	color = "R"

	img = loadImg(inImage)
	keyPts = NextKeyPoints(inImage, cfgEdge, "")
	draw(img, keyPts, rs, color, outImage)

	cfgEdge = "./cfg/BC-edge.json"
	color = "G"

	inImage = outImage
	img = loadImg(inImage)
	keyPts = NextKeyPoints(inImage, cfgEdge, "")
	draw(img, keyPts, rs, color, inImage)

	fmt.Println()
}
