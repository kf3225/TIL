DI(Dependency Injection) = 依存性注入
==

依存関係分離
--

## 問題のある処理の流れ

1. mainからの呼び出し<br>
⬇️
2. マルチプレクサによる関数呼び出し<br>
⬇️
3. 各HTTPメソッドのHandle関数<br>
⬇️
4. DBから値を取ってくる関数<br>
⬇️
5. グローバル変数sql.DB<br>
⬇️
6. PostgreSQL

とあるとすると、2は3に、3は4に、4は5に、とそれぞれ依存関係にある（＝それが無いと処理として成立しない関係）

原因は**sql.DB**　➡️　こいつへの依存を除去できないか？

- sql.DBの使用をいかにして絶つか**という考えでは無い**
    - 直接の依存をいかにして避けるか
- トップダウンで考え最初から依存原因となっているものを注入してしまう
    - 1の時点で2を呼び出すときに注入する

<br>

## 解決策

1. mainからの呼び出し<br>
⬇️　⬅️ sql.DBを注入
2. マルチプレクサによる関数呼び出し<br>
⬇️　⬅️ sql.DBを注入
3. 各HTTPメソッドのHandle関数<br>
⬇️　⬅️ sql.DBを注入
4. DBから値を取ってくる関数<br>
⬇️
5. DIされたsql.DB<br>
⬇️
6. PostgreSQL

<br>

## ハードル&解決

**hundleFuncの第二引数はレスポンスとリクエストを引数とするシグネチャの関数でなければならない!!**<br>
⬇️<br>
handlerFunc型となるような関数を返却する関数を規定してあげればいい

```go

// Text interface
type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

// この構造体にはポインタレシーバでfetch、create、update、deleteが実装されている
// つまりダックタイピングからTextインターフェースだと見做せる
// この構造体がsql.DBのポインタを持っている
type Post struct {
	Db      *sql.DB
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	var err error
	db, err := sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr: "127.0.0.1:8080",
    }
    // HundleFuncの第二引数にはレスポンス、リクエストを引数としてとる関数を返す関数が設定
    // さらに上記の関数は*sql.DBをフィールドに持つ構造体を引数にとる
    // 構造体にはmainで最初に作ったdbの接続情報を設定する
	http.HandleFunc("/post/", handleRequest(&Post{Db: db}))

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequest(t Text) http.HandlerFunc {
    // 関数を返却
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case http.MethodGet:
			err = handleGet(w, r, t)
		case http.MethodPost:
			err = handlePost(w, r, t)
		case http.MethodPut:
			err = handlePut(w, r, t)
		case http.MethodDelete:
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
```

構造体のフィールドに*sql.DBを持たせる事で直接handleRequestに*sql.DBを渡さなくても大丈夫になった

<br>

## DIの目的

単体テストをやりやすくするため

```go
package main

// FakePostはTextインターフェースを実装する構造体
type FakePost struct {
	ID      int
	Content string
	Author  string
}

// そのままIDを返却
func (post *FakePost) fetch(id int) (err error) {
	post.ID = id
	return
}

func (post *FakePost) create() (err error) {
	return
}

func (post *FakePost) update() (err error) {
	return
}

func (post *FakePost) delete() (err error) {
	return
}

```

handleGetに対するテストの例だと下記のようになる

```go
func setup() {
    mux = http.NewServeMux()
    // ここでPostの代わりに上のFakePostを渡す！
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Fatalf("Response code is %d", writer.Code)
	}

	actual, err := decodeTestData(writer)
	if err != nil {
		log.Fatal(err)
	}
	expected := 1

	if expected != actual.ID {
		t.Errorf("expected : %d, but was : %d", expected, actual.ID)
	}
}
```

このようにデータベースに依存しないテストができる

# ただしfetchのテストを行なっていない事には留意しなければならない