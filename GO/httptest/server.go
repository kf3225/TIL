package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

// Person struct
type Person struct {
	ID        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Age       int     `json:"age"`
	Sex       int     `json:"sex"`
	Address   Address `json:"address"`
}

// Address struct
type Address struct {
	PersonID    int    `json:"personId"`
	Prefecture  string `json:"prefecture"`
	Town        string `json:"town"`
	HouseNumber string `json:"houseNumber"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	if _, err := os.Stat("person"); os.IsNotExist(err) {
		err = os.Mkdir("person", 0755)
		if err != nil {
			log.Println(err)
			return
		}
	}

	http.HandleFunc("/person/", handleRequest)

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
		log.Println(err)
	}

	return
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println(err)
		return
	}

	fileDir := "./person"
	fileName := fmt.Sprintf("person_%05d.json", id)
	if _, err = os.Stat(path.Join(fileDir, fileName)); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	fileData, err := os.Open(path.Join(fileDir, fileName))
	if err != nil {
		log.Println(err)
		return
	}
	defer fileData.Close()

	b, err := ioutil.ReadAll(fileData)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := findLastNum()
	if err != nil {
		log.Println(err)
		return
	}
	id = id + 1
	fileDir := "person"
	fileName := fmt.Sprintf("person_%05d.json", id)
	jsonFile, err := os.Create(path.Join(fileDir, fileName))
	if err != nil {
		log.Println(err)
		return
	}
	defer jsonFile.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var person Person
	if err = json.Unmarshal(b, &person); err != nil {
		log.Println(err)
		return
	}

	person.ID = id
	person.Address.PersonID = id
	if b, err = json.MarshalIndent(person, "", "\t"); err != nil {
		log.Println(err)
		return
	}

	if _, err = jsonFile.Write(b); err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)

	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println(err)
		return
	}

	fileDir := "./person"
	fileName := fmt.Sprintf("person_%05d.json", id)

	if _, err = os.Stat(path.Join(fileDir, fileName)); os.IsNotExist(err) {
		log.Println(err)
		return
	}
	jsonFile, err := os.OpenFile(path.Join(fileDir, fileName), os.O_WRONLY, 0755)
	if err != nil {
		log.Println(err)
		return
	}
	defer jsonFile.Close()

	var person Person
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r.Body)
	if err = json.Unmarshal(buffer.Bytes(), &person); err != nil {
		log.Println(err)
		return
	}

	person.ID = id
	person.Address.PersonID = id
	b, err := json.MarshalIndent(&person, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	if _, err = jsonFile.Write(b); err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	return

}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println(err)
		return
	}
	fileDir := "./person"
	fileName := fmt.Sprintf("person_%05d.json", id)

	if _, err = os.Stat(path.Join(fileDir, fileName)); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	if err = os.Remove(path.Join(fileDir, fileName)); err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func findLastNum() (lastNum int, err error) {
	files, err := ioutil.ReadDir("./person")
	if err != nil {
		log.Println(err)
		return
	}

	var max int
	for _, file := range files {
		if lastNum, err = strconv.Atoi(file.Name()[7:12]); err != nil {
			log.Println(err)
			return
		}

		if lastNum > max {
			max = lastNum
		}
	}
	return
}
