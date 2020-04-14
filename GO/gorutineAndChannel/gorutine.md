gorutineとchannel
==

並行処理と並列処理
--

- 並行処理<br>
    多くのものを同時に**扱う**(ex. 2つの行列、1つのレジ)<br>
    線で表したときにひとつに交わる＝実行時間に重なりがある

    ➡️　データを共有したり実行のタイミングを合わせたり（そこまでをそれぞれで実行。単一リソースでも可能）

- 並列処理<br>
    多くのものを同時に**行う**(ex. 2つの行列、2つのレジ)<br>
    線で表したときに平行のまま

    ➡️　タスクを同時実行。複数リソース（CPUなど）を要する

GOはどちらも書ける（並行かつ並列なども可能）

<br>

ゴルーチン
--

複数が同時実行される関数　≠　スレッド（）

関数にgoをつけるだけ（名前付でも無名でも）

### WaitGroup

次の処理を開始する前に全てのゴルーチンが終わっているかを確認

終わってない場合はブロック

1. WaitGroupを宣言する
2. Addを使用してWaitGroupのカウンタを初期化する
3. ゴルーチンが作業完了する度に、Doneを使ってカウンタを減産
4. Waitを呼び出すとカウンタが0になるまで次の処理の実行をブロックする

```go 
func main() {
	var wg sync.WaitGroup // WaitGroupを宣言
	wg.Add(2) // カウンタの初期化（2つのゴルーチンなので2）
	go printLetter(&wg) // 引数にWaitGroupのポインタを渡す
	go printNumber(&wg) // 引数にWaitGroupのポインタを渡す
	wg.Wait() // 各関数内でDoneによりカウンタが2→0まで減算されるのを待つ

	fmt.Println("end") // カウンタが0になったあとようやく実行
}

func printNumber(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
    wg.Done() // 関数を抜ける前にDoneでカウンタ減算
}

func printLetter(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	wg.Done() // 関数を抜ける前にDoneでカウンタ減算
}
```

***ゴルーチンの減算をし忘れるとdeadlock！***<br>
➡️　WaitGroupは全てのゴルーチンがスリープしていることをランタイムが検出するまで実行を中断し、検出されるとエラーになる