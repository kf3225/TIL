Ginkgo
==

## Ginkgoとは

- BDDスタイル＝振る舞い（Behavior）駆動開発
- BDDはTDD（テスト駆動開発）の拡張版
- テスト手法ではなく**開発手法**
- ユーザーストーリーを作る（業務上の要件を集積したもの）

## ユーザーストーリー

```
Story : "投稿の取得"
In order to "投稿をユーザーに提示する"
As a "呼び出し側プログラム"
I want to "投稿を取得する"

Scenario 1 : "idを使う"
Given "投稿のidが1"
When "そのidのGETリクエストを送信した"
Then "投稿を取得する"

Scenario 2 : "非整数のidを使う"
Given "投稿のidがHelloWorld"
When "そのidのGETリクエストを送信した"
Then "HTTP 500のレスポンスを取得する"
```

## 実際のテストの書き方

```shell
ginkgo bootstrap
ginkgo generate
```

テストの雛形を自動生成する

自動生成されたxxx_test.goのDescribe関数の中に書いていく

```go
BeforeEach(func() {
		post = &FakePost{}
		mux = http.NewServeMux()
		writer = httptest.NewRecorder()

		mux.HandleFunc("/post/", HandleRequest(post))
	})
```
各テストケースの前に実行される処理を記述する場合はBeforeEach関数

<br>

```go
// ケースの説明
Context("Get a post using an id", func() {
        // It ➡️　そのケースがどうなるべきか
		It("should get a post", func() {
			request, _ := http.NewRequest(http.MethodGet, "/post/1", nil)
			mux.ServeHTTP(writer, request)

            // gomegaのマッチャーで期待値と実際の値を比較
			Expect(writer.Code).To(Equal(http.StatusOK))

			var post Post
			json.Unmarshal(writer.Body.Bytes(), &post)

            // gomegaのマッチャーで期待値と実際の値を比較
			Expect(post.ID).To(Equal(1))
		})
	})
```

```shell
ginkgo -v
```
上記コマンドでテスト実行