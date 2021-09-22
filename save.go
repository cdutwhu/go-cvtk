package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func saveJPG(img image.Image, path string) image.Image {
	out, err := os.Create(path)
	if err != nil {
		return nil
	}
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 100
	if err := jpeg.Encode(out, img, &opts); err != nil {
		log.Println(err)
	}
	return img
}

func savePNG(img image.Image, path string) image.Image {
	out, err := os.Create(path)
	if err != nil {
		return nil
	}
	defer out.Close()

	if err := png.Encode(out, img); err != nil {
		log.Println(err)
	}
	return img
}
