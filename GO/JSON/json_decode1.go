package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Post struct
type Post struct {
	ID int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}

// Author struct
type Author struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

// Comment struct
type Comment struct {
	ID int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post_r.json")
	if err != nil {
		fmt.Println("Error opening JSON File :", err)
		return
	}
	defer jsonFile.Close()

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON File :", err)
		return
	}

	var post Post
	err = json.Unmarshal(b, &post)
	if err != nil {
		fmt.Println("Error unmarshaling JSON File :", err)
	}
	fmt.Println(post)
}