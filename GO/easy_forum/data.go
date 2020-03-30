package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	dbInfo := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", "gwp", "gwp", "gwp", "disable")
	Db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Error connecting DB :", err)
	}
}

func Retrive(id int) (post Post, err error) {
	post := Post{}
	row := Db.QueryRow("SELECT id, content, author FROM post WHERE id = $1", id).Scan(&post.ID, &post.Content, &post.author)
	return
}

func Create(post *Post) (err error) {
	statement := "INSERT INTO post (id, content, author) VALUES ($1, $2, $3) RETURNING id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.ID, post.Content, post.Author).Scan(&post.ID)
	return
}

func Update(post *Post) (err error) {
	updatestr := "UPDATE post SET content = $2, author = $3 WHERE id = $1"
	_, err := Db.Exec(updatestr, post.ID, post.Content, post.Author)
	return
}

func Delete(post *Post) (err error) {
	deletestr := "DELETE FROM post WHERE id = $1"
	_, err := Db.Exec(deletestr, post.ID)
	return
}
