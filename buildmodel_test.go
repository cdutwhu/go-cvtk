package main

import (
	"image/color"
	"testing"
)

func TestBuildModel(t *testing.T) {

	BuildModel("./in/mel.png", color.RGBA{255, 0, 0, 255})

}
