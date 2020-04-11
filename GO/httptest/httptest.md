Golangのテスト
==========

テスト用の標準ライブラリは以下の2つある
- testing
- httptest ⬅️ 今回つかうやつ

httptest
----------

1. ## httptestの基本

    1. マルチプレクサを生成
    2. テスト対象のハンドラをマルチプレクサに付加
    3. レコーダを作成（writer）
    4. リクエストを作成（request）
    5. テスト対象のハンドラにリクエストを送信し、レコーダに記録
    6. レコーダより結果をチェック

<br>

2. ## httptestを使用した単体テスト

```go
func TestHandleDelete(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/person/", handleRequest)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/person/1", json) // ⬅️　読み込んだJSONファイルなど。io.Reader
	mux.ServeHTTP(writer, request) // レコーダーとリクエストができたらServeHTTPに渡してマルチプレクサに送信

	if writer.Code != http.StatusOK { // ⬅️　テストケースによるステータスコード
		t.Fatalf("Response code is %d", writer.Code)
	}

    // ここらへんでJSONを構造体にシリアライズしたりする処理
    // writer.Body.Bytes()を実際値の構造体へ
    
    expected := "期待値" // ⬅️　期待値を設定

    if expected != actual.Value {
        t.Errorf("expected %s, but was %s", expected, actual.Value)
    }
}
```

注意点として
- 各テストケースは独立して実行　➡️　それぞれでテスト用Webサーバを立ち上げるため各関数内で上記の1〜6の処理が必要！
- マルチプレクサはレスポンスをブラウザに送信する代わりにレコーダーに渡すようになっている

<br>

## httptestのテクニック

各テストケース実行前に入れたい処理がある場合は以下のように書く

```go
// mux & writerをグローバル変数化
var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) { 
	setup()         // セットアップ関数が起動
    code := m.Run() // テストケースが実行（１つ）
    // teardown()   // テスト終了後の処理を書きたければここへ
    os.Exit(code)   // 終了（次のテストケースがあればまたセットアップから）
}

// setup関数にmuxとwriterを初期化処理を書く
func setup() {
	mux = http.NewServeMux()
	mux.HandleFunc("/person/", handleRequest)
	writer = httptest.NewRecorder()
}
```