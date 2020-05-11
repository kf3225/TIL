Redux-Saga
==

外部APIを利用する＝JavaScript（TypeScript）のため、***非同期処理である***<br />
通信等の副作用を含む非同期通信処理をReduxのmiddlewareでやろう

そのmiddleware = Redux-Saga

## 部品

- select : Store Stateから必要な値を取得する
- put : Action Creatorを実行してActionをDispatchする
- take : 特定のActionを待ち受ける（take = takeEvery, takeLatest）
  - takeEvery : Actionがdispatchされるたびにredux-sagaのタスクが起動する -> ex) F5連打された際のリクエスト分実行
  - takeLatest : Actionがdispatchされるたびにredux-sagaのタスクが起動する。すでに同じActionによって起動したタスクがまだ終了していない場合は、そのタスクはキャンセルされる -> ex) F5連打された際のリクエストの最後だけ実行
- call : 外部の非同期処理関数をコールする
- fork : 自分とは別のスレッドを起動し、そこで特定のタスクを実行する。Taskオブジェクトを返す
- join : forkの戻り値のTaskオブジェクトを指定して、そのタスクが完了するのを待つ

https://redux-saga.js.org/docs/api/

