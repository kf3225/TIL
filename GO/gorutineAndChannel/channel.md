Channel
==

チャネルとは
--

ゴルーチン同士を相互通信させる方法<br>
➡️　箱のようなもの

チャネルはスライスなどと同様makeによって割り当てる<br>
➡️　得られる値は**実際に使用されるデータ構造への参照**

```go
ch1 := make(chan int)       // バッファなし
ch2 := make(chan int, 10)   // バッファあり
```

### 1. バッファなしチャネル<br>
    同期的に扱われる＝一つの箱のようなもの

    とあるゴルーチンが箱に物を入れる<br>
    ⬇️<br>
    他のゴルーチンがものを入れようと思っても物が入ってるの入れられない（箱が空になるまでスリープ）<br>
    ⬇️<br>
    別のゴルーチンが箱に入ってる物を取る<br>
    ⬇️<br>
    他のゴルーチンが箱から物を取り出そうとしても物がないので取り出せない（箱に物が入るまでスリープ）<br>

```go
ch1 := make(chan <-string) // 受信専用チャネル
ch1 <- "A" // 物を入れたらスリープが解除され次の処理へ
// 次の何かしらの処理

ch2 := make(<-chan string) // 送信専用チャネル
str := <-ch2 // 物を入れたらスリープが解除され次の処理へ
// 次の何かしらの処理
```

```go
func thrower(c chan int) {
	for i := 0; i < 5; i++ {
        // 引数のチャネルがバッファなしなのでここの処理で数値を入れた後、
        // このチャネルから入れた数値が取り出されるまで次の数値を入れられずにスリープ
		c <- i
		fmt.Println("Threw >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
        // チャネルに数値が入ってくるまでスリープ
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

func main() {
    // バッファなしチャネルを宣言
	c := make(chan int)
	go thrower(c)
	go catcher(c)

	time.Sleep(10 * time.Millisecond)
}
```

実行結果

```
Threw >> 0
Caught << 0
Caught << 1
Threw >> 1
Threw >> 2
Caught << 2
Caught << 3
Threw >> 3
Threw >> 4
Caught << 4
```

- ThrewよりCaughtが先に来てる数値があるがこれはfmt.Printlnの実行順序がゴルーチンによって前後したにすぎない
- 0のスロー・キャッチから4のスロー・キャッチまで数値がちゃんと昇順になっている<br>
    = **現在の数値をキャッチする前に次の数値がスローされないようにブロックされている**

### 2.バッファありチャネル

バッファなし＝同期的<br>
バッファあり＝非同期的、先入れ先出し（FIFO）

箱の例だとゴルーチンは箱に空きがなくなるまで入れ続けることができ、他のゴルーチン**入ってきた順番で**取り出し続けることができる<br>
取り出すものがなくなったときに作業中断

バッファなしチャネルのソースをバッファありで宣言する

```go
func main() {
    // バッファが3のチャネルを宣言
	c := make(chan int, 3)
	go thrower(c)
	go catcher(c)

	time.Sleep(10 * time.Millisecond)
}
```

実行結果

```
threw >> 0
caught << 0
caught << 1
threw >> 1
threw >> 2
threw >> 3
threw >> 4
caught << 2
caught << 3
caught << 4
```

バッファなしと違い非同期実行となっている

バッファ付チャネルは業務を処理するプロセスに制限がある時に便利<br>
➡️　プロセスに渡す要求の数を絞れる

<br>

select
--

selectはチャネル版switch文のようなもの<br>

- Aケースのチャネルに値が入っていて、Bケースには入っていない : Aケースが選択
- Bケースのチャネルに値が入っていて、Aケースには入っていない : Bケースが選択
- どちらのケースにも値がある : Goのランタイムがどちらかをランダムに選択
- どちらのケースにも値がない : defaultブロックに入る

※defaultブロックがない場合デッドロックになる（ならないようにする方法もある＝close）<br>
	➡️　A, B両方のゴルーチンから値を取り出され、両方のゴルーチンがスリープする為


```go
func callerA(c chan string) {
	c <- "Hello World!"
	// close関数で文字列受信後にチャネルを閉じた状態にする
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	// close関数で文字列受信後にチャネルを閉じた状態にする
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)

	go callerA(a)
	go callerB(b)

	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {

		select {
		// msgに文字列を送信、ok1にチャネルがcloseされたかのbool値
		// callerAでチャネルに文字列が送信されたあともチャネルがcloseされるまでtrue
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A, ok1 is %v now\n", msg, ok1)
			}
			fmt.Println("ok1", ok1)
		// msgに文字列を送信、ok1にチャネルがcloseされたかのbool値
		// callerBでチャネルに文字列が送信されたあともチャネルがcloseされるまでtrue
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B, ok2 is %v now\n", msg, ok2)
			}
			fmt.Println("ok2", ok2)

		}
	}
}
```

- チャネルがcloseしてもそのチャネルが使えなくなるということではない
- チャネルcloseはそのチャネルにもう何も値が送信されないことを受信側に示すだけ