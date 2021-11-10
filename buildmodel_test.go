package main

import (
	"image/color"
	"testing"
)

func TestBuildModel(t *testing.T) {

	BuildModel("./in/start/1.jpg", "./out/AB-edge.json", color.RGBA{255, 0, 0, 255})

}
