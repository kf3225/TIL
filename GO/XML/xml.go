package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Post struct
type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	// XMLStr   string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

// Comment struct
type Comment struct {
	ID      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

// Author struct
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens", err)
			return
		}

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
	}

}
