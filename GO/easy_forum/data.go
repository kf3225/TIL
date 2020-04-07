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

func RetriveCommentByCommentId(commentId int) (post Post, err error) {
	sql := `
	SELECT
	    p.id,
	    p.content,
	    pa.id,
	    pa.name,
	    c.c_id,
	    c.c_content,
	    c.c_author_id,
	    c.c_author_name
	FROM post as p
	INNER JOIN author as pa
	ON p.author_id = pa.id
	INNER JOIN (
	    SELECT
	        cm.id       as c_id,
	        cm.content  as c_content,
	        cm.post_id  as post_id,
	        ca.id       as c_author_id,
	        ca.name     as c_author_name
	    FROM comment as cm
	    INNER JOIN author as ca
	    ON cm.author_id = ca.id
	) as c
	ON p.id = c.post_id
	AND p.delete_flag = '0'
	WHERE c.c_id = $1
	ORDER BY c.c_id ASC
	;
	`
	var comment Comment
	err = Db.QueryRow(sql, commentId).Scan(&post.ID, &post.Content, &post.Author.ID, &post.Author.Name,
		&comment.ID, &comment.Content, &comment.Author.ID, &comment.Author.Name)
	if err != nil {
		fmt.Print("Error scanning Query :", err)
		return
	}
	post.Comments = append(post.Comments, comment)
	return
}

func RetrivePostById(id int) (post Post, err error) {

	sql := `
	SELECT
	    p.id,
	    p.content,
	    pa.id,
	    pa.name,
	    c.c_id,
	    c.c_content,
	    c.c_author_id,
	    c.c_author_name
	FROM post as p
	INNER JOIN author as pa
	ON p.author_id = pa.id
	INNER JOIN (
	    SELECT
	        cm.id       as c_id,
	        cm.content  as c_content,
	        cm.post_id  as post_id,
	        ca.id       as c_author_id,
	        ca.name     as c_author_name
	    FROM comment as cm
	    INNER JOIN author as ca
	    ON cm.author_id = ca.id
	) as c
	ON p.id = c.post_id
	WHERE p.id = $1
	AND p.delete_flag = '0'
	ORDER BY c.c_id ASC
	;
	`
	rows, err := Db.Query(sql, id)
	if err != nil {
		fmt.Println("Error query :", err)
	}
	defer rows.Close()

	for rows.Next() {
		comment := Comment{}

		err = rows.Scan(&post.ID, &post.Content, &post.Author.ID, &post.Author.Name,
			&comment.ID, &comment.Content, &comment.Author.ID, &comment.Author.Name)
		if err != nil {
			println("Error scanning :", err)
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	return
}

func (p *Post) createPost() (err error) {
	statement := "INSERT INTO post (id, content, author) VALUES ($1, $2, $3) RETURNING id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(p.ID, p.Content, p.Author).Scan(&p.ID)
	return
}

func (p *Comment) createComment() (err error) {
	stmt := "INSERT INTO comment (content, author_id) VALUES ($1, $2)"
	statement, err := Db.Prepare(stmt)
	if err != nil {
		fmt.Println("Error Prepare statement :", err)
	}
	defer statement.Close()
	return
}

func (p *Post) update(err error) {
	updatestr := "UPDATE post SET content = $2, author = $3 WHERE id = $1"
	_, err = Db.Exec(updatestr, p.ID, p.Content, p.Author)
	return
}

func (p *Post) deletePost() (err error) {

	if p.ID == 0 {
		return
	}

	sqlPost := "UPDATE post SET delete_flg = '1' WHERE id = $1 RETURNING id;"
	sqlComment := "UPDATE comment SET delete_flg = '1' WHERE post_id = $1; RETURNING id"
	var id int
	err = Db.QueryRow(sqlPost, p.Comments[0].ID).Scan(&id)
	if err != nil {
		fmt.Println("Error update post :", err)
		return
	}

	err = Db.QueryRow(sqlComment, p.ID).Scan(&id)

	return
}

func (p *Post) deleteComment() (err error) {
	sql := "UPDATE comment SET delete_flg = '1' WHERE id = $1 RETURNING id"
	var id int
	err = Db.QueryRow(sql, p.Comments[0].ID).Scan(&id)
	return
}
