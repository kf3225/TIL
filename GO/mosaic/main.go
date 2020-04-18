package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

func main() {
	// マルチプレクサを生成
	mux := http.NewServeMux()

	// 静的ファイルのルートディレクトリを指定(この外へのアクセスは不可)
	// FileServerに渡すことによって静的ファイルを返すハンドラを生成
	files := http.FileServer(http.Dir("public"))
	// filesを元に先頭の/static/を除外するハンドラを生成
	staticFilesHandle := http.StripPrefix("/static/", files)
	// URL/static/と生成した静的ファイルハンドラを紐付け
	// HTMLのCSSやJavaScriptの参照先に/static/指定した時に
	// pubulic内の静的ファイルを見に行ってくれる！
	mux.Handle("/static/", staticFilesHandle)

	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	TILESDB = tileDB()
	fmt.Println("Mosaic Server started")

	server.ListenAndServe()

}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	// maxMemoryを指定
	r.ParseMultipartForm(32 << 20)
	// imageの要素からファイルを取得
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))

	// 取得したファイルを画像にデコード
	original, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 取得した画像から画像情報を生成
	bounds := original.Bounds()

	// モザイク加工後の画像の型枠を生成
	newimage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))

	db := cloneTilesDB()

	sp := image.Point{0, 0}

	// 画面で選択したタイルサイズを1マスとして元画像サイズ分処理を行なっていく
	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {

			// pixel値を取得
			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}

			// TILLESDBから近似色のタイル画像ファイル名を取得
			nearest := findNearest(color, &db)
			file, err := os.Open(nearest)
			if err == nil {
				img, _, err := image.Decode(file)
				if err == nil {
					// TILESDBから取得した画像を指定したサイズでリサイズ
					t := resize(img, tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)

					draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
				} else {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()
		}
	}
	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newimage, nil)
	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())

	t1 := time.Now()

	images := map[string]string{
		"original": originalStr,
		"mosaic":   mosaic,
		"duration": fmt.Sprintf("%v", t1.Sub(t0)),
	}
	t, _ := template.ParseFiles("result.html")

	t.Execute(w, images)
}
