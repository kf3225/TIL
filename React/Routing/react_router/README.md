# React Router

使用モジュール

- react-router-dom
- react-router
- @types/react-router-dom
- @types/react-router

## react-router-dom

1. \<BrowserRouter>

HTML5 の HistoryAPI が使用可能になる<br />
https://developer.mozilla.org/ja/docs/Web/API/History_API

src/index.tsx

```typescript
import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import App from "./App";
import * as serviceWorker from "./serviceWorker";
import { BrowserRouter } from "react-router-dom";

ReactDOM.render(
  <BrowserRouter>
    <App />
  </BrowserRouter>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
```

<br />

2. \<Link>

```typescript
<ListItem
  button
  component={Link}
  to={`characters/${characterData[code].school}`}
>
```

Link コンポーネントの呼び出し。to=URL のリンクを生成

\<a>タグを使用すると react-router の管轄外となる<br />
Web サーバーにリクエストが飛んで SPA のページ全体がリロードされる

<br />

## react-router

1. \<Switch>, \<Route>, \<Redirect>

src/App.tsx

```typescript
import * as React from "react";
import { Redirect, Route, Switch } from "react-router-dom";
import "./App.css";

import Characters from "./components/Characters";
import Home from "./components/Home";

const App: React.FC<{}> = () => (
  <div className="container">
    <Switch>
      <Route path="/characters/:code" component={Characters} />
      <Route path="/" component={Home} />
      <Redirect to="/" />
    </Switch>
  </div>
);

export default App;
```

- \<Switch> -> コンポーネント毎にルーティング可能
- \<Route> -> path に URL がマッチすると component にした Component がレンダリング
- \<Redirect> -> どのエントリにもマッチしなかった場合の処理

<br />

2. withRouter

src/components/Characters/index.tsx

```typescript
import { parse } from "query-string";
import * as React from "react";
import { RouteComponentProps, withRouter } from "react-router";
import { Redirect } from "react-router-dom";
import { Button } from "@material-ui/core";
import { Home } from "@material-ui/icons";

import { characterData } from "../../characterData";
import Spinner from "../common/Spinner";
import { Helmet } from "react-helmet";
import CharacterList from "./CharacterList";
import "./index.css";

type CharacterListProps = {} & RouteComponentProps<{ code: string }>;
```

src/App.tsx で`<Route path="/characters/:code" component={Characters} />`とあり、
この URL の`:code`部分を RouteComponentProps の型引数として取得可能（常に string 型）

今回 Character で Props を使用しないので空のオブジェクトと RouteComponentProps を&で繋いで合成している<br />
Props を使用する場合はいつも通り interface で Props を宣言して合成すれば OK

<br />

src/components/Characters/index.tsx

```typescript
const Characters: React.FC<CharacterListProps> = ({
  history,
  location,
  match,
}) => {
  const codes = Object.keys(characterData);
  const targetCode = match.params.code;
  const isLoading = parse(location.search).loading === "true";
```

query-string は URL からクエリパラメータを抽出するモジュール<br />
location オブジェクトから loading という URL のパラメータを抜き出す

<br />

src/components/Characters/index.tsx

```typescript

  return codes.includes(targetCode) ? (
    <>
      <Helmet>
        <title>Characters</title>
      </Helmet>
      <header>
        <h1>CharacterList / {characterData[targetCode].school}</h1>
      </header>
      {isLoading ? (
        <Spinner />
      ) : (
        <CharacterList key={characterData[targetCode].school}
          school={characterData[targetCode].school}
          characters={characterData[targetCode].players}
        />
      )}
      <Button onClick={() => history.push("/")}>
        <Home />
        Home
      </Button>
    </>
  ) : (
    <Redirect to="/" />
  );
};

export default withRouter(Characters);
```

export するコンポーネントを withRouter という HOC で包むことによって関数コンポーネントの引数として下記 Props が使用可能に

- history
  - push() -> そのURLにいた履歴が残る
  - replace() -> そのURLにいた履歴が抹消される

- location
  - pathname -> 置換後パスからクエリを抜いた部分。上記ケースだと`/characters/AAA`
  - search -> クエリ部分。`?loading=true`など

- match
  - params(objects) -> key/value形式。上記ケースだと`{code: "AAA"}`など
  - isExact(boolean) -> 後続に文字もなく、マッチしている場合はtrue
  - path -> 置換前パスからクエリを抜いた部分。上記ケースだと`/characters/:code`
  - url -> 置換後パスからクエリを抜いた部分。上記ケースだと`/characters/AAA`
