package main

func main() {
	// defer gotk.TrackTime(time.Now())

	// f, err := os.Open("./test.jpg")
	// if err != nil {
	// 	// Handle error
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// img, fmtName, err := image.Decode(f)
	// if err != nil {
	// 	// Handle error
	// 	log.Fatalln(err)
	// }
	// fmt.Println(fmtName)
	// imagerect := img.Bounds()
	// fmt.Println(img.Bounds())

	// // fmt.Println(img.ColorModel())

	// left, top, right, bottom := imagerect.Min.X, imagerect.Min.Y, imagerect.Max.X, imagerect.Max.Y

	// roi := ROIgray(img, left, top, right, bottom)

	// // roi := ROIgrayV2(img, 200, 200, 100)
	// // roi := ROIrgbaV2(img, 99, 99, 120)

	// // for i, p := range roi.Pix {
	// // 	fmt.Printf("%d ", p)
	// // 	if i > 100 {
	// // 		break
	// // 	}
	// // }

	// hImg := DrawBlob(left, top, right, bottom, roi.Pix)
	// save(hImg, "./out_blob.jpg")
}
