package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
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
	PostID  int    `json:"post_id"`
	Author  Author `json:"author"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case http.MethodGet:
		err = handleGet(w, r)
	case http.MethodPost:
		err = handlePost(w, r)
	case http.MethodPut:
		err = handlePut(w, r)
	case http.MethodDelete:
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handling http Get Method
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	var id int
	var post Post
	for key, value := range r.URL.Query() {
		for _, v := range value {
			id, err = strconv.Atoi(v)

			if err != nil {
				fmt.Println("Error casting str to int :", err)
				return
			}

			switch key {
			case "post":
				if post, err = RetrivePostById(id); err != nil {
					fmt.Println("Error Retriving Post data by post id :", err)
					return
				}
			case "comment":
				if post, err = RetriveCommentByCommentId(id); err != nil {
					fmt.Println("Error Retriving Post data by commtn id :", err)
					return
				}
			}
		}
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshaling JSON :", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// // Handling http Post method
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {

	switch filepath.Base(r.URL.Path) {
	case "post":
		var post Post
		if post, err = decodePostJson(r); err != nil && err != io.EOF {
			fmt.Println("Error decode Post json :", err)
			return
		}
		err = (&post).createPost()
	case "comment":
		var comment Comment
		if comment, err = decodeCommentJson(r); err != nil && err != io.EOF {
			fmt.Println("Error decode Comment json :", err)
			return
		}
		err = (&comment).createComment()
	}

	if err != nil {
		fmt.Println("Error create data :", err)
		return
	}

	return
}

// // Handling http Put method
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {

	var id int
	var post Post
	var comment Comment
	var author Author
	for key, value := range r.URL.Query() {
		for _, v := range value {
			if id, err = strconv.Atoi(v); err != nil {
				fmt.Println("Error cast string to int :", err)
				return
			}
			switch key {
			case "post":
				if post, err = decodePostJson(r); err != nil && err != io.EOF {
					fmt.Println("Error decode post Json :", err)
					return
				}
				err = (&post).updatePost(id)
			case "comment":
				if comment, err = decodeCommentJson(r); err != nil && err != io.EOF {
					fmt.Println("Error decode comment Json :", err)
					return
				}
				err = (&comment).updateComment(id)
			case "author":
				if author, err = decodeAuthorJson(r); err != nil && err != io.EOF {
					fmt.Println("Error decode author Json :", err)
					return
				}
				err = (&author).updateAuthor(id)
			}
		}
	}
	if err != nil {
		fmt.Println("Error update data :", err)
		return
	}
	return
}

// Handling http Delete method
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	var id int
	var post Post
	for key, value := range r.URL.Query() {
		for _, v := range value {
			if id, err = strconv.Atoi(v); err != nil {
				fmt.Println("Error casting str to int :", err)
				return
			}
			switch key {
			case "comment_id":
				if post, err = RetriveCommentByCommentId(id); err != nil {
					fmt.Println("Error retrive comment by comment id :", err)
					return
				}
				if err = (&post).deleteComment(); err != nil {
					fmt.Println("Error delete comment :", err)
					return
				}
			case "post_id":
				if post, err = RetrivePostById(id); err != nil {
					fmt.Println("Error retrive post by post id :", err)
					return
				}
				if err = (&post).deletePost(); err != nil {
					fmt.Println("Error delete post :", err)
					return
				}
			}
		}
	}

	output, err := json.MarshalIndent(&post, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func decodeCommentJson(r *http.Request) (comment Comment, err error) {
	decoder := json.NewDecoder(r.Body)

	for {
		err = decoder.Decode(&comment)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error decode json :", err)
			return
		}
	}
	return
}

func decodePostJson(r *http.Request) (post Post, err error) {
	decoder := json.NewDecoder(r.Body)

	for {
		err = decoder.Decode(&post)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error decode json :", err)
			return
		}
	}
	return
}

func decodeAuthorJson(r *http.Request) (author Author, err error) {
	decoder := json.NewDecoder(r.Body)

	for {
		err = decoder.Decode(&author)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error decode json :", err)
			return
		}
	}
	return
}
