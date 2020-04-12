package main

import (
	"database/sql"
)

// Text interface
type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

// Post struct
type Post struct {
	Db      *sql.DB
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (p *Post) fetch(id int) (err error) {
	s := "SELECT id, content, author FROM post WHERE id = $1"
	err = p.Db.QueryRow(s, id).Scan(&p.ID, &p.Content, &p.Author)
	return
}

func (p *Post) create() (err error) {
	s := "INSERT INTO post (content, author) VALUES ($1, $2) RETURNING id"
	err = p.Db.QueryRow(s, p.Content, p.Author).Scan(&p.ID)
	return
}

func (p *Post) update() (err error) {
	s := "UPDATE post SET content = $2, author = $3 WHERE id = $1 RETURNING id"
	err = p.Db.QueryRow(s, p.ID, p.Content, p.Author).Scan(&p.ID)
	return
}

func (p *Post) delete() (err error) {
	s := "DELETE FROM post WHERE id = $1 RETURNING id"
	err = p.Db.QueryRow(s, p.ID).Scan(&p.ID)
	return
}
