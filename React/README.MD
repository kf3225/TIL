React and TypeScript
==

## TypeScriptメモ

**1. ReturnType**
```typescript
const foo = (params: Params): Hoge => ({hoge: params});
type hoge = ReturnType<typeof foo>; // Hoge型
```

<br />

**2. as const**
```typescript
const foo = "TYPE_FOO";
const bar = "TYPE_BAR" as "TYPE_BAR"; // 可変な変数に代入した時の型変化予防策(従来)
const hoge = {
  foo, // -> string型 : 可変な変数に代入すると型が変化してしまう
  bar, // -> TYPE_BAR型
}
```

```typescript
const foo = "TYPE_FOO" as const; // 可変な変数に代入した時の型変化予防策(<- new!)
const bar = "TYPE_BAR" as "TYPE_BAR"; // 可変な変数に代入した時の型変化予防策(従来)
const hoge = {
  foo, // -> TYPE_FOO型
  bar, // -> TYPE_BAR型
}
```

```typescript
const hoge = {
  foo: "TYPE_FOO", // -> TYPE_FOO型
  bar: "TYPE_BAR", // -> TYPE_BAR型
} as const // readonly属性付与
```

<br />

**3. Generics**

```typescript
interface foo {
  foo1: string;
  foo2: number;
}
const hoge = <T, U, P extends foo>(arg1: T, arg2: U, arg3: P) => ({
  a: arg1,
  b: arg2,
  c: arg3, // args3はinterface fooの要素を満たしていなければコンパイルエラー
});
```

<br />

## axiosメモ

>Promise based HTTP client for the browser and node.js

https://github.com/axios/axios

<br />

**1. create()**

```typescript
// AxiosRequestConfigを生成
const instance = axios.create({ 
  baseURL: `https://some-domain.com/api/`,
  timeout: 7000,
  headers: { "Content-Type": "application/json" },
  responseType: 'json',
});
```
