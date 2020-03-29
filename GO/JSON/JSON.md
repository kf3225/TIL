JSON(Decode)
============

<br>

## Unmarshal & Decode 使い分け

入力によって使い分ける
| 使い分け   | 入力パターン          | 入力例　                   |
|-----------|---------------------|--------------------------|
| Unmarshal | 文字列データやメモリ内  | 文字列データ、メモリ内データ |
| Decode    | io.Readerのストリーム | http.RequestのBody       |

<br>

## JSONを用意

```json
{
    "id": 1,
    "content": "Hello World",
    "Author": {
        "id": 2,
        "name": "Keisuke"
    },
    "comments":[
        {
            "id": 3,
            "content": "Have a great day!",
            "author": {
                "id": 4,
                "name": "Taro"
            }
        },
        {
            "id": 5,
            "content": "How are you today?",
            "author": {
                "id": 6,
                "name": "Jiro"
            }
        }
    ]
}
```

<br>

## 構造体を定義

```go
type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}
```

```go
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
```

```go
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  Author `json:author`
}
```

<br>

## Unmershalの場合

```go
jsonFile, err := os.Open("post_r.json")
```
ファイルをオープン
```go
b, err := ioutil.ReadAll(jsonFile)
```
ファイルをEOFになるまで読み込む(Unmarshalを試したいので[]byteを返すReadAll使用)
```go
var post Post
err = json.Unmarshal(b, &post)
```
json.Unmarshalにバイト配列と構造体の参照を渡して対応付させる

<br>

実行結果
```
{1 Hello World {2 Keisuke} [{3 Have a great day! {4 Taro}} {5 How are you today? {6 Jiro}}]}
```

<br>

## Decodeの場合
```go
jsonFile, err := os.Open("post_r.json")
```
ファイルをオープンする

Decodeを使いたいのでos.File（io.Readerをインターフェースに持つ)を返すOpenの戻り値を利用
```go
decoder := json.NewDecoder(jsonFile)
```
デコーダーを生成する
```go
var post Post
	for {
		err = decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON File :", err)
			return
		}
	}
```
EOFになるまで無限ループ

生成したデコーダーに構造体の参照を渡してDecodeすると構造体にデータが格納される

<br>

実行結果
```
{1 Hello World {2 Keisuke} [{3 Have a great day! {4 Taro}} {5 How are you today? {6 Jiro}}]}
```

<br>

## メモ

#### io.Reader　-> 文字列([]byte)を読み出すためのインターフェース
- os.Openの戻り値の*os.Fileはio.Readerインターフェースを持ってる！
- http.RequestのBodyメンバはReadCloserというインターフェースでReaderメソッドが実装されている

#### io.Readerの中身
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
このシグネチャを持っているメソッドが実装されている型はio.Readerに該当

-> これが世に言うダックタイピングか？