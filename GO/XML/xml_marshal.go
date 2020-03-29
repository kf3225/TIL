package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Post struct
type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

// Author struct
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		ID:      "1",
		Content: "Hello World",
		Author: Author{
			ID:   "1",
			Name: "Keisuke",
		},
	}

	buffer := &bytes.Buffer{}
	_, err := buffer.Write([]byte(xml.Header))
	if err != nil {
		fmt.Println("Error wirte XML Header to XML :", err)
		return
	}

	encoder := xml.NewEncoder(buffer)

	encoder.Indent("", "\t")

	err = encoder.Encode(&post)

	err = ioutil.WriteFile("post_w.xml", []byte(buffer.String()), 0644)

}
