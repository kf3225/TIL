# Hooks

Local State をクラスを使わずに関数コンポーネントで書ける！

```typescript
import React, { FC, useState } from "react";
```

インポートで useState を呼び出す

```typescript
const [count, setCount] = useState(0);
```

useState()に初期値を渡すと戻り値としてその初期値とセッターを返すので分割代入で受け取る

複数の Local State を使いたい時はその分書き連ねて宣言するのがセオリー

```typescript
const [count, setCount] = useState(0);
const [str, setStr] = useState("str");
const [bool, setBool] = useState(true);
```

<br />

# Effect Hook

副作用を扱う Hooks でライフサイクルメソッド＝ componentDidMount(), componentDidUpdate(), componentWillUnmount()を扱いたい場合に使用

※副作用 -> ログ、データ取得、手動での DOM の改変 etc...

```typescript
useEffect(() => { // 第一引数に引数なしの関数を設定

  // componentDidMount(), componentDidUpdate()メソッド内に書くのと同義
  // effect hookでは上記2つがまとめられている
  doSomething();

  // 戻り値なしも可。戻り値ありの場合はcomponentWillUnmount()メソッド内に処理を書いているのと同義
  return clearSomething();

}, [watchVar]); // 第二引数省略可。あれば配列で
```

1. 初回レンダリング直後にdoSomething()が呼ばれる(componentDidMount()と同じ)
2. 再レンダリング時にwatchVarが変化していればまたdoSomethin()が呼ばれる(componentDidUpdate()と同じ)
3. n回目レンダリング時にwatchVarに変化が無い場合doSomething()は呼ばれず、アンマウント直前でclearSomething()が呼ばれる(componentWillUnmount()と同じ)

※第二引数watchVarを省略or空配列にした場合は初回レンダリング時にdoSomething()が呼ばれてそれ以降は実行されない
※useEffect()は一つの関数コンポーネントで複数回呼べるため、やりたい処理によって書き分けるのが吉
