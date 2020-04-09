Golangのテスト
==========

テスト用の標準ライブラリには以下の2つある
- testing
- net/http/httptest

<br>

## testing

### 1. testingの基本
- *_test.goの形にする
- go testで呼び出す(実行ディレクトリ内の*_test.goが全て実行)
- go test -v (verbose)で詳細なテスト結果表示
- go test -v -coverでカバレッジ取得
- テスト対象と同一パッケージ内に配置する

<br>

### 2. testingを使用した単体テスト

```go
func TestXxx(t *testing.T) {
    expected := ... // 期待値を生成する処理
    actual := Xxx(...) // 実際にテスト対象を利用して値を生成する処理

    // 同値性の確認（値が一緒だったらOK）
    if expected != actual {
        t.Errorf("expected : %+v, but was : %+v", expected, actual)
    }

    // 同値性の確認（値が一緒だったらOK）
    if reflect.deepEqual(expected, actual) {
        t.Errorf("expected : %+v, but was : %+v", expected, actual)
    }

    // 同一性の確認（参照しているアドレスが一緒だったらOK）
    if &expected != actual {
        t.Errorf("expected : %+v, but was : %+v", &expected, &actual)
    }
}

func TestYyy(t *testing.T) {
    // このテスト対象はスキップ
    t.Skip("Skipping Yyy for now")
}
```

<br>

### 3. testing.Tで使用できるメソッド
- Log     : fmt.Printに似た関数。エラーログ生成
- Logf    : fmt.Printfに似た関数。フォーマット書式子等利用してエラーログ生成
- Fail    : テスト関数に失敗した記録を残し、実行を継続する
- FailNow : テスト関数に失敗した記録を残し、実行を停止する

上記を組み合わせた簡易関数

|         |  Log   |  Logf  |              |
|---------|:------:|:--------:|:-----------|
|  Fail   |  Error | Errorf | ⬅️停止しない   |
| FailNow |  Fatal | Fatalf | ⬅️停止する     |
|         |    ⬆️  |    ⬆️   |              |
|         |  Print | Printf |              |

<br>

### 4. testingテクニック

#### Skip
```go
func TestLongRunningTest(t *testing.T) {
    if testing.Short() {
        t.Skip("Skip test")
    }
    // わざと10秒止める
    time.Sleep(10 * time.Second)
}
```

```shell
go test -v -short
```

10秒の処理がスキップされる

#### parallel

```go
func TestParallel_1(t *testing.T) {
    t.Parallel()
    time.Sleep(2 * time.Second)
}

func TestParallel_2(t *testing.T) {
    t.Parallel()
    time.Sleep(4 * time.Second)
}

func TestParallel_3(t *testing.T) {
    t.Parallel()
    time.Sleep(6 * time.Second)
}
```

```shell
go test -v -short -parallel 3
```
3つのテストケースが平行実行される

<br>

### 5. ベンチマーク

```go
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("test.json")
	}
}
```

```shell
go test -v -short -bench .
```

テスト対象のベンチマークを測る事ができる(すべてのベンチマークを実行するため-benchに.=ドットを渡す)

```shell
go test -v -test.count 5 -short -bench .
```

-test.count [0-9]+を指定することによってテストの実行回数指定可能

```shell
go test -v -test.count 5 -short -run x -bench .
```
-runはどの機能テストを実行するかのフラグ。xを指定しているがそんなテスト関数は無いので機能テストは全て無視されてベンチマークだけがcount分走る
