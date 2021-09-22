package main

import "image"

func SplitRGBA(img image.Image) (r, g, b, a *image.Gray) {
	rect := img.Bounds()
	rgba := ROIrgba(img, rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
	r, g, b, a = image.NewGray(rect), image.NewGray(rect), image.NewGray(rect), image.NewGray(rect)

	for i, p := range rgba.Pix {
		switch i % 4 {
		case 0:
			r.Pix[i/4] = p
		case 1:
			g.Pix[i/4] = p
		case 2:
			b.Pix[i/4] = p
		case 3:
			a.Pix[i/4] = p
		}
	}

	// wg := &sync.WaitGroup{}
	// wg.Add(4)
	// go func(rgbaBytes, chBytes []byte) {
	// 	for i, j := 0, 0; i < len(rgbaBytes); i += 4 {
	// 		chBytes[j] = rgbaBytes[i]
	// 		j++
	// 	}
	// 	wg.Done()
	// }(rgba.Pix[0:], r.Pix)
	// go func(rgbaBytes, chBytes []byte) {
	// 	for i, j := 0, 0; i < len(rgbaBytes); i += 4 {
	// 		chBytes[j] = rgbaBytes[i]
	// 		j++
	// 	}
	// 	wg.Done()
	// }(rgba.Pix[1:], g.Pix)
	// go func(rgbaBytes, chBytes []byte) {
	// 	for i, j := 0, 0; i < len(rgbaBytes); i += 4 {
	// 		chBytes[j] = rgbaBytes[i]
	// 		j++
	// 	}
	// 	wg.Done()
	// }(rgba.Pix[2:], b.Pix)
	// go func(rgbaBytes, chBytes []byte) {
	// 	for i, j := 0, 0; i < len(rgbaBytes); i += 4 {
	// 		chBytes[j] = rgbaBytes[i]
	// 		j++
	// 	}
	// 	wg.Done()
	// }(rgba.Pix[3:], a.Pix)
	// wg.Wait()

	return
}

// func SplitHSV(img image.Image) (h, s, v *image.Gray) {
// 	// rect := img.Bounds()
// 	// rgba := ROIrgba(img, rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
// 	// h, s, v = image.NewGray(rect), image.NewGray(rect), image.NewGray(rect)
// }
