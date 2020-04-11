package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strconv"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	mux = http.NewServeMux()
	mux.HandleFunc("/person/", handleRequest)
	writer = httptest.NewRecorder()

	if err := os.RemoveAll("person"); err != nil {
		log.Fatal(err)
	}
	log.Println("Remove all file in person dir")

	os.Mkdir("person", 0755)
	log.Println("make dir person")

}

func prepareData() {
	r := regexp.MustCompile(`test\_data[0-9]\.json`)
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		if r.MatchString(file.Name()) {
			id, _ := strconv.Atoi(file.Name()[9:10])
			cpFileName := fmt.Sprintf("person/person_%05d.json", id)
			testData, err := os.Open(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer testData.Close()
			dstFile, err := os.Create(cpFileName)
			if err != nil {
				log.Fatal(err)
			}
			defer dstFile.Close()

			b, err := ioutil.ReadAll(testData)
			if err != nil {
				log.Fatal(err)
			}
			var person Person
			json.Unmarshal(b, &person)
			person.ID = id
			person.Address.PersonID = id
			b, _ = json.MarshalIndent(&person, "", "\t")

			if _, err = dstFile.Write(b); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func TestHandlePost(t *testing.T) {

	body1, _ := os.Open("test_data1.json")

	request, err := http.NewRequest(http.MethodPost, "/person/", body1)
	if err != nil {
		t.Fatalf("person Request Body : %+v, %v", body1, err)
	}

	mux.ServeHTTP(writer, request)
	if writer.Code != http.StatusOK {
		t.Fatalf("Request Response status : %v", writer.Code)
	}
	if _, err := os.Stat("./person/"); os.IsNotExist(err) {
		t.Error("expected making json file", err)
	}
}

func TestHandleGet(t *testing.T) {

	prepareData()

	request, _ := http.NewRequest(http.MethodGet, "/person/1", nil)
	mux.ServeHTTP(writer, request)

	log.Println(request.URL.Path)
	if writer.Code != http.StatusOK {
		os.RemoveAll("person")
		t.Fatalf("Response code is %v", writer.Code)
	}

	var person Person
	if err := json.Unmarshal(writer.Body.Bytes(), &person); err != nil {
		os.RemoveAll("person")
		t.Fatal("json unmarshal error : ", err)
	}

	var expected int = 1
	if person.ID != expected {
		t.Errorf("expected : %d, but was : %d", expected, person.ID)
	}
	os.RemoveAll("person")

}

func TestHandlePut(t *testing.T) {

	prepareData()

	b := []byte(`
	{
		"id": 0,
		"firstName": "Update",
		"lastName": "Suzuki",
		"age": 20,
		"sex": 1,
		"address": {
			"personId": 0,
			"prefecture": "Tokyo",
			"town": "Shinjuku",
			"houseNumber": "CCC 1-2-3"
		}
	}
	`)
	body := bytes.NewBuffer(b)

	request, _ := http.NewRequest(http.MethodPut, "/person/2", body)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		os.RemoveAll("person")
		t.Fatalf("Response code is %v", writer.Code)
	}

	var actual Person
	actualFile, err := os.OpenFile("person/person_00002.json", os.O_RDONLY, 0755)
	if err != nil {
		t.Fatal("open error :", err)
	}
	defer actualFile.Close()
	ab, _ := ioutil.ReadAll(actualFile)

	if err = json.Unmarshal(ab, &actual); err != nil {
		log.Fatal(err)
	}

	expected := "Update"
	if actual.FirstName != expected {
		t.Errorf("expected : %s, but was : %s", expected, actual.FirstName)
	}
	os.RemoveAll("person")

}

func TestHandleDelete(t *testing.T) {

	prepareData()

	request, _ := http.NewRequest(http.MethodDelete, "/person/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Fatalf("Response code is %d", writer.Code)
	}

	if _, err := os.Stat("person_00001.json"); os.IsExist(err) {
		t.Error("file did not delete")
	}
	os.RemoveAll("person")
}
