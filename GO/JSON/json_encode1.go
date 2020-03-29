package main

import (
	"encoding/json"
	"fmt"
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
	Author  Author `json:"author"`
	Content string `json:"content"`
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
				ID: 3,
				Author: Author{
					ID:   4,
					Name: "Taro",
				},
				Content: "Have a great day!",
			},
			Comment{
				ID: 5,
				Author: Author{
					ID:   6,
					Name: "Jiro",
				},
				Content: "How are you today?",
			},
		},
	}
	jsonFile, err := os.Create("post_w1.json")
	if err != nil {
		fmt.Println("Error creating JSON File :", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)

	if err != nil {
		fmt.Println("Error encoding JSON File :", err)
		return
	}
}
