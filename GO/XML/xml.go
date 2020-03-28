package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Post struct
type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	XMLStr  string   `xml:",innerxml"`
}

// Author struct
type Author struct {
	ID   string `xml:id,attr`
	Name string `xml:"chardata"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file", err)
		return
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post.XMLName)
	fmt.Println(post.ID)
	fmt.Println(post.Content)
	fmt.Println(post.Author)
	fmt.Println(post.XMLStr)
}
