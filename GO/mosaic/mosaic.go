package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"os"
	"path"
)

// TILESDB メモリ展開グローバル変数
var TILESDB map[string][3]float64

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
				i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}
	return *out
}

// タイル画像データをメモリ展開
func tileDB() map[string][3]float64 {
	fmt.Println("start populating tiles db...")

	db := make(map[string][3]float64)
	readDir := "tiles"
	files, _ := ioutil.ReadDir(readDir)
	for _, f := range files {
		name := path.Join(readDir, f.Name())
		file, err := os.Open(name)

		if err == nil {
			img, _, err := image.Decode(file)

			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating TILE DB:", err, name)
			}
		} else {
			fmt.Println("cannot open file:", err, name)
		}
		file.Close()
	}
	fmt.Println("finish populating tiles db")

	return db
}

// メモリ展開したタイルデータをクローン
func cloneTilesDB() map[string][3]float64 {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	return db
}

// 色の近似値ファイルを探す
func findNearest(target [3]float64, db *map[string][3]float64) string {
	var fileName string

	smallest := 1000000.0

	for k, v := range *db {
		dist := distance(target, v)
		// 前の周の距離より短ければ距離とファイル名を保存
		// 最終的に残った物が近似値ファイル
		if dist < smallest {
			fileName, smallest = k, dist
		}
	}
	delete(*db, fileName)
	return fileName
}

// 距離を算出
// 2次元におけるユークリッド距離 : √(x1-y1)^2+(x2-y2)^2
// n次元におけるユークリッド距離 : √(x1-y1)^2+(x2-y2)^2 ... (xn-yn)^2
func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}

func sq(n float64) float64 {
	return n * n
}
