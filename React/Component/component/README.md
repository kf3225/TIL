Propsメモ
==

## import

```typescript
import CharacterList, { Character } from './CharacterList';
```

1. default exportされた"CharacterList"クラスをインポート
2. default exportされた"CharacterList"クラスから"Character"インターフェースをインポート

## Props

タグの属性値として表現されるもの＝親コンポーネントから受け取る値のこと

- クラスコンポーネント
  - そのクラスのメンバ変数
- 関数コンポーネント
  - その関数の引数

```typescript
export interface Character {
  id: number;
  name: string;
  age: number;
  height?: number;
}

interface CharacterListProps {
  school: string;
  characters: Character[];
}
```

↑　まずはインターフェース

```typescript
class CharacterList extends Component<CharacterListProps> {
```

genericsでPropsの型を"CharacterListProps"に指定<br />
指定しない場合は（Propsを使う必要のない時）は何も指定しない<br />
何も指定しないと"{}"＝空オブジェクトが設定されている状態

このコンポーネントをタグとしてマウントする時に必要な属性値とその型が決定される<br />
  →　上記の場合は"CharacterList"というタグに属性が"school"(string型), "characters"(Characterオブジェクト配列型)となる

```typescript
const { school, characters } = this.props;
```

"CharacterList"タグの属性として渡された"school", "characters"の値をPropsのメンバ変数としてアクセス可能

```typescript
<Item.Group>
  {characters.map(c => (
  /* eslint-disable react/jsx-key */
    <Item key={c.id}>
      <Icon name="user circle" size="huge" />
      <Item.Content>
        <Item.Header>name : {c.name}</Item.Header>
        <Item.Meta>age    : {c.age}</Item.Meta>
        <Item.Meta>height : {c.height ? c.height : '???'} cm</Item.Meta>
      </Item.Content>
    </Item>
  /* eslint-enable */
  ))}
</Item.Group>
```

\<Item key={c.id}>は\<Item>でも動くが**ワーニング**(Itemはreact-semantic-uiのコンポーネント)<br />
→ reactでループ処理を行う際は一意になるキー属性値を設定しないと警告になる。<br />
→ 最小限の変更でDOMに反映する時のパフォーマンスを効率化するため



