コンポーネントライフサイクル
==

### 1. マウント : コンポーネントが呼ばれてDOMに挿入されること。メソッドが以下の順で呼ばれる
  - constructor() : コンストラクタ<br />
  ↓
  - static getDerivedStateFromProps() : レンダリングの直前で呼ばれて、戻り値でLocal Stateを変更可能<br />
  ↓
  - render() : レンダリングを行う<br />
  ↓
  - ***componentDidMount() : コンポーネントがマウントされた直後に呼ばれる***<br />

<br />

### 2. アップデート : propsやstateに変更時に以下のメソッドが呼ばれる
  - static getDerivedStateFromProps() : レンダリングの直前で呼ばれて、戻り値でLocal Stateを変更可能<br />
  ↓
  - ***shouldComponentUpdate() : 再レンダリングの直前に呼ばれてfalseを返せば再レンダリングを中止可能***<br />
  ↓
  - render() : レンダリングを行う<br />
  ↓
  - getSnapshotBeforeUpdate() : コンポーネントが変更される直前に呼ばれて戻り値でスナップショットを取っておける<br />
  ↓
  - ***componentDidUpdate() : コンポーネントが変更された直後に呼ばれる***<br />

<br />

### 3. アンマウント : コンポーネントがDOMから削除される時に以下のメソッドが呼ばれる
  - ***componentWillUnmount() : コンポーネントがアンマウントされる直前に呼ばれる***

<br />

### 4. エラーハンドリング : 子コンポーネントのレンダー中、ライフサイクルメソッド内、またはコンストラクタ内でエラーが発生した時に以下のメソッドが呼ばれる
  - static getDerivedStateFromError() : 子孫コンポーネントで例外が起きた時に呼ばれてStateを更新する<br />
  ↓
  - componentDidCatch() : 子孫コンポーネントで例外が起きた時に呼ばれる

<br />

## 今回のケース

```typescript
componentDidMount = () => {
  this.timerId = setInterval(this.tick, 1000);
}
```

マウント直後にsetInterval関数によって一秒毎に延々とtick関数を呼び続ける処理が起動される<br />

<br />

```typescript
componentDidUpdate = () => {
  const { timeLeft } = this.state;
  if (timeLeft === 0) {
    this.reset();
  }
}
```

setIntervalから呼ばれるtickによってLocal Stateがずっと一秒ずつ更新されるのでその度にcomponentDidUpdate関数が呼ばれる<br />
LocalStateのtimeLeftの値が0になったらreset関数を呼び、60からリスタートされる

<br />

```typescript
componentWillUnmount = () => {
  clearInterval(this.timerId as NodeJS.Timer);
}
```

setIntervalはタスクが生き続けてしまうのでcomponentDidMount時に取得したidでタスクキルする処理が必要<br />
アンマウントされる直前にその処理を挟む
