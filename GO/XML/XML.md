GOのXMLの取り扱い
============

## XMLを定義

```xml
<?xml version="1.0" encoding="utf-8" ?>
<post id="1">
    <content>Hello World!</content>
    <author id="1">Keisuke</author>
    <comments>
        <comment id="1">
            <content>Have a good day!</content>
            <author id="2">Taro</author>
        </comment>
        <comment id="2">
            <content>How are you today?</content>
            <author id="3">Jiro</author>
        </comment>
    </comments>
</post>
```

<br>

## 構造体を定義

```go
type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	XMLStr   string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}
```

```go
type Comment struct {
	ID      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}
```

```go
type Author struct {
	ID   string `xml:id,attr`
	Name string `xml:",chardata"`
}
```

<br>

## 構造体への格納

```go
decoder := xml.NewDecoder(xmlFile)
```
オープンしたファイルからデコーダー作成

```go
t, err := decoder.Token()
```
デコーダーからトークン（＝XML要素を表すインターフェース）を取得
```go
switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "comment":
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				//fmt.Println(comment)
			case "author":
				var author Author
				decoder.DecodeElement(&author, &se)
				//fmt.Println(author)
			case "post":
				var post Post
				decoder.DecodeElement(&post, &se)
				fmt.Println(post)
			}
		}
```
トークンの型チェックしてXMLデータを構造体に格納していく

StartElementは開始タグのこと

<br>

## 実行結果

```
{{ post} 1 Hello World! {1 Keisuke} [{1 Have a good day! {2 Taro}} {2 How are you today? {3 Jiro}}]}
```

<br>

## 構造体とXML要素の対応付が必要
  - ``で囲まれた部分は構造体タグ
  - key : valueの形をとる文字列は""で囲む
    
      → 構造体タグが``で囲まれているから""を使用
  - 対応付はGo言語が行うため**全ての場所から参照可能なフィールドにする**

      → **大文字から始める**

<br>

## 忘れそう
  1. 要素名 : xml.Nameを使用
  2. 要素の属性   : \`xml:<属性名>,attr\`
  3. 文字データ   : XML要素タグと同名フィールドに\`xml:,chardata\`
  4. 未処理データ : \`xml:,innerxml\`
  5. モードデータ（,attr・,chardata・,innerxmlなど）がない場合は構造体のフィールドは構造体名と同じ名前のXMLの要素と紐付けられる
  6. 木構造を使ってでなく要素を直接取得したい場合は\`xml:"a>b>c"\`を使用する（aとbは中間要素でcが取得したいノード）

