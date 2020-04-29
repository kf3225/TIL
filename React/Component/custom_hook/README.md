# Custom Hook

## コンポーネントを役割で二つに分割

1. コンポーネント : Presentational Component -> 表示部分
2. コンテナ : Container Component -> Local State を含む Props などの処理部分

|            | component                                      | container                                             |
| ---------- | ---------------------------------------------- | ----------------------------------------------------- |
| 関心       | どのように見えるか                             | どのように機能するか                                  |
| 送受信     | データや振る舞いを Props として一方的に受信    | データや振る舞いを他のコンポーネントに送信する        |
| Flux       | Flux の Store 等に依存しない                   | Flux の Action 実行や Flux の Store に依存            |
| データ変更 | データ変更に介入しない                         | データ変更に介入し、任意の処理を行う                  |
| 書き方     | 関数コンポーネントが多数                       | HOC や Render Props（オワコンらしい）や Hooks（熱い） |
| Dir        | ./components/AAA.tsx<br />./components/BBB.tsx | ./containers/AAA.tsx<br />./containers/BBB.tsx        |

<br />

## 今回のケース

### component

```typescript
interface AppProps {
  timeLeft: number;
  keep: boolean;
  reset: () => void;
  stopOrRestart: () => void;
}

const AppComponent: FC<AppProps> = ({ timeLeft, keep, reset, stopOrRestart }) => (
```

ボタンクリック時の関数や Local State を外から注入する形で、それに合わせて interface を定義する

<br />

### container

**_useTimer 関数が今回の肝の Custom Hooks(関数名は useXxx とするのが慣習。しないと Eslint に怒られる)_**

```typescript
const useTimer = (
  limitSec: number
): [number, boolean, () => void, () => void] => {
  const [timeLeft, setTimeLeft] = useState(limitSec);
  const [keep, setKeep] = useState(true);
```

Local State に必要なカウントダウンの値とスタートストップ状態、reset 関数と stopOrRestart 関数を配列に設定した戻り値を返す

引数から取得したカウントとスタートストップ状態の初期値で useState()から state と setter を取得する

初期値を引数から取得する形にしているので 1 分や 5 分などのカウントダウンタイマーとして対応可能

```typescript
// ボタンクリック時のreset関数
const reset = () => {
  setTimeLeft(limitSec);
};

// ボタンクリック時のstopOrRestart関数
const stopOrRestart = () => {
  setKeep(!keep);
};

useEffect(() => {
  const tick = () => {
    setTimeLeft((prevTime) => (prevTime === 0 ? limitSec : prevTime - 1));
  };
  let timerId: NodeJS.Timer;
  if (keep) {
    timerId = setInterval(tick, 1000);
  }
  return () => clearInterval(timerId);
}, [keep, limitSec]);

return [timeLeft, keep, reset, stopOrRestart];
};
```

useEffect()でマウント時や更新時の挙動やアンマウント時の挙動を規定している

useEffect()の第二引数の値に変化がある場合、<br />
例えばlimitSecが変わって1分から2分のカウントダウンタイマーへ変わると値の変化を察知して<br/>
useEffect()が再実行され、tick()の挙動が再定義されることによってカウントが0まで進んだ後の<br />
リセット後の値が1分から2分へと変更される
