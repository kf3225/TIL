Local State
==

## state

```typescript
interface AppState {
  count: number;
}

class App extends Component<{}, AppState> {
  constructor(props: {}) {
    super(props);
    this.state = { count: 0};
  }
```

Componentのジェネリクスとして空のオブジェクトをまず呼ぶ -> 今回はPropsを使わないため<br />
constructor引数として空オブジェクトであるpropsを渡す
コンストラクタの中で最初にsuper(props)を呼んでいる -> JavaScript(TypeScript)言語の性質上、<br />
thisを使用できるようにするには**親クラスのコンストラクタを呼ばないといけないため、super呼び出しは絶対！**

stateをAppStateインターフェースの型と同様に初期化count=0<br />
stateは**コンストラクタ内のみで直接設定可能（this.state = hogeとできるのはコンストラクタ内のみ）**

<br />

## stateへの値の設定

```typescript
increment() {
  this.setState(prevState => ({
    count: prevState.count + 1,
  }));
}

decrement() {
  this.setState(prevState => ({
    count: prevState.count - 1,
  }));
}
```

stateに値を設定するには上記のように専用のセッターを使用する<br />
参照は普通の呼び出しで可能

- setStateに設定できる引数

1. 設定したい値のStateオブジェクト -> Stateに固定値を設定するケース
2. setState(prevState, props) => newState形式のState&Propsを引数として受け取って新しいStateを返す関数（上記の例ではPropsを使わないので省略している） -> Stateを動的に変更するケース

<br />

```typescript
render() {
  const {count} = this.state;
  return (
    <div className="container">
      <header><h1>カウンター</h1></header>
      <Card>
        <Statistic className="number-board">
          <Statistic.Label>count</Statistic.Label>
          <Statistic.Value>{count}</Statistic.Value>
        </Statistic>
        <Card.Content>
          <div className="ui two buttons">
            <Button color="red" onClick={() => this.decrement()}>-1</Button>
            <Button color="green" onClick={() => this.increment()}>+1</Button>
          </div>
        </Card.Content>
      </Card>
    </div>
  );
}
```

***親コンポーネントの値を子コンポーネントで書き換えたい時は自身の状態を変化させる関数を子コンポーネントに渡してイベントで発火するようにその関数を仕込んでおく***

今回の場合はAppコンポーネントのStateのメンバであるcountの値をButtonコンポーネント内で書き換えたいのでcountの状態を変化させるincrementとdecrement関数をonClickイベントで発火するように仕込んでいる！！！←考え方が大事

<br />

() => this.decrement()はthis.decrementに書き換えられない？ -> 不可能。この場合のthisは実行時のオブジェクト＝Buttonコンポーネントになるため（ButtonコンポーネントにはsetState()なんてないから）

this.decrementにする場合はdecrement関数をアロー関数に書き換える必要あり

```typescript
decrement = (e: SyntheticEvent): void =>  {
  e.preventDefault();
  this.setState(prevState => ({
    count: prevState.count - 1,
  }));
}
```

上記のSyntheticEventはイベントハンドラのオブジェクト。preventDefault()でonClickの挙動を制御している -> \<a>タグなどの場合は画面遷移してしまうがそれをしたくない場合はこのような記述になる（今回はButtonのため特に設定する必要はない）
