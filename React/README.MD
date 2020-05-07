React and TypeScript
==

## TypeScriptメモ

1. ReturnType
```typescript
const foo = (params: Params): Hoge => ({hoge: params});
type hoge = ReturnType\<typeof foo>; // Hoge型
```
