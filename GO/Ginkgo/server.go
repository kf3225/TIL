package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	var err error
	db, err := sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", HandleRequest(&Post{Db: db}))

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func HandleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case http.MethodGet:
			err = handleGet(w, r, t)
		case http.MethodPost:
			err = handlePost(w, r, t)
		case http.MethodPut:
			err = handlePut(w, r, t)
		case http.MethodDelete:
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println(err)
		return
	}

	if err = post.fetch(id); err != nil {
		log.Println(err)
		return
	}

	b := encode(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

	return
}

func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {

	if err = decode(r, post); err != nil {
		log.Fatal(err)
	}
	err = post.create()

	return
}

func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {

	if err = decode(r, post); err != nil {
		log.Fatal(err)
	}
	err = post.update()
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err = post.fetch(id); err != nil {
		log.Fatal(err)
	}
	err = post.delete()
	return
}

func decode(r *http.Request, post Text) (err error) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	buffer := bytes.NewBuffer(body)
	decoder := json.NewDecoder(buffer)

	for {
		err := decoder.Decode(&post)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

func encode(post Text) []byte {

	var body bytes.Buffer
	encoder := json.NewEncoder(&body)

	encoder.SetIndent("", "\t")
	encoder.Encode(&post)

	return body.Bytes()
}
