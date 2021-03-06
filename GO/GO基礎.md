GOの積み上げ
============
GO言語とかWebの基礎とかメモ  


RESTful API
------------
  - RESTベースのWebサービスはリソースに注目してHTTPのメソッドがリソースに働きかける
  - RESTの考え方・発展 : オブジェクト指向！
    - データを表す「**モデル**」（**オブジェクト**と呼ばれる）
    - モデルに付属する「**関数**」（**メソッド**と呼ばれる）
  - REST = **リソースと呼ばれるモデル**に対して少数のアクション（HTTPメソッド）を許可
  - HTTP経由の場合はURLでリソースを表し、HTTPメソッドが動詞  
    | HTTPメソッド | 目的                              | 例               |
    |-----------|-----------------------------------|------------------|
    | POST      | リソースの生成（リソースが存在しない場合） | POST /users      |
    | GET       | リソースの取得                       | GET /users/1     |
    | PUT       | 指定されたURLによるリソースの更新       | PUT /users/1     |
    | DELETE    | リソースの削除                       | DELETE /users/1  |
    | PATCH     | リソースの部分的更新                 | PATCH /users/1   |
    
      - PUT : どのリソースが置換されるか正確にしっておく必要あり。与えられたURLにひとつのリソースが作成される
      - POST : 与えられたURLを呼び出すたびに新リソースが作成される
  - HTTPメソッドの使用目的は決まっている → 複雑なアクション（例ユーザーアカウントのアクティベート）はどうすれば？
    1. アクションを名詞化してリソースに
        1. ユーザー有効状態を表す有効化リソース作成
        2. 有効化リソースに追加属性を持たせることも（発展性）
      ```http:HTTP POST
        POST /users/456/activattion HTTP/1.1
        {"date": "2020-01-01-T23:59:59Z"}
      ```
        
    2. アクションをリソースの属性に
       1. ユーザー有効状態は単にリソースの属性と定義
       2. リソースの一部分の更新のためPATCHを使う
      ```http:HTTP POST
        PATCH /users/456/activattion HTTP/1.1
        {"active": true}
      ```
        
