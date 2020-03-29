package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Post struct
type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

// Author struct
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Comment struct
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  Author `json:author`
}

func main() {
	jsonFile, err := os.Open("post_r.json")
	if err != nil {
		fmt.Println("Error opening Json File :", err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	var post Post
	for {
		err = decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON File :", err)
			return
		}
	}
	fmt.Println(post)
}
