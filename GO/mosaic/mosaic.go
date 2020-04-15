package main

import (
	"image"
	"image/color"
)

// 画像（区分けされたセクションを想定）の平均色を算出
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()

	r, g, b := 0.0, 0.0, 0.0

	// 取得した画像のX&Y軸の初期座標から最大まで
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			r1, g1, b1, _ := img.At(x, y).RGBA()

			// 各値をMAX.X * MAX.Y回のループ分溜め込む
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}

	totalPixels := float64(bounds.Max.X * bounds.Max.Y)

	// 溜め込んだ各値を総ピクセル数で割って平均を算出
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

// リサイズ関数
func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()

	ratio := bounds.Dx() / newWidth

	// 何も書き込まれていないリサイズ後の画像を生成する
	out := image.NewNRGBA(
		image.Rect(bounds.Min.X/ratio, bounds.Min.Y/ratio, bounds.Max.X/ratio, bounds.Max.Y/ratio))

	// 取得した画像のX&Y軸の初期座標から最大まで
	// x, yは画像を何分の1にするかというratioの逆数分進む
	// i, jはリサイズ後画像の座標
	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(
				i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> b), uint8(a >> 8)})
		}
	}
	return *out
}
