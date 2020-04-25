yarn
==

パッケージモジュール管理ツール

よく使うのは
- yarn add \<PACKAGE_NAME\> : PACKAGE_NAMEをインストール。package.json参照
- yarn remove \<PACKAGE_NAME\> : PACKAGE_NAMEをアンインストール
- yarn upgrade \<PACKAGE_NAME\> : PACKAGE_NAMEを最新に更新
- yarn info \<PACKAGE_NAME\> : PACKAGE_NAMEの情報表示

<br>

React始める前のECMAScript メモ
==

## 高階関数=Higher Order Function
```javascript
const hof = (ex, fn) => n => fn(n + ex);
const plusOneDouble = hof(1, n => n * 2);
console.log(plusOneDouble(4)); //  = 10
```

```
> const hof = (ex, fn) => n => fn(n + ex);
undefined
> const plusOneDouble = hof(1, n => n * 2);
undefined
> console.log(plusOneDouble(4));
10
undefined
```

コンポーネントを引数にとってコンポーネントを返却するHigher Order Component(HOC)というテクニックに使用

<br>

## クロージャ=Closuer

```javascript
const counterMaker = (initialCount) => {
    let c = initialCount;
    const increment = () => c++;
    return increment;
};
const count = counterMaker(1);
console.log(count(), count(), count());
```

```
> const counterMaker = (initialCount) => {
...     let c = initialCount;
...     const increment = () => c++;
...     return increment;
... };
undefined
> const count = counterMaker(1);
undefined
> console.log(count(), count(), count());
1 2 3
undefined
```

クロージャ＝ローカル変数を参照する関数を戻り値とした関数<br>
（通常ローカル変数は関数処理終了後に破棄される）

<br>

## カリー化

```javascript
const multi = n => m => n * m;
console.log(multi(2)(5));
const triple = multi(3);
console.log(triple(5));
```

```
> const multi = n => m => n * m;
undefined
> console.log(multi(2)(5));
10
undefined
> const triple = multi(3);
undefined
> console.log(triple(5));
15
undefined
```

カリー化=値を返す関数をネストした関数<br>
関数の部分適用が可能 => 関数の一部の引数を固定して新しい関数を作ること