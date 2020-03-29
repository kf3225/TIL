package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Author  Author `json:"author"`
}

func main() {
	post := Post{
		ID:      1,
		Content: "Hello World!",
		Author: Author{
			ID:   2,
			Name: "Keisuke",
		},
		Comments: []Comment{
			Comment{
				ID:      3,
				Content: "Have a great day!",
				Author: Author{
					ID:   4,
					Name: "Taro",
				},
			},
			Comment{
				ID:      5,
				Content: "How are you today?",
				Author: Author{
					ID:   6,
					Name: "Jiro",
				},
			},
		},
	}

	b, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshaling JSON File :", err)
		return
	}
	err = ioutil.WriteFile("post_w2.json", b, 0644)
}
