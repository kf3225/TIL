package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var writer *httptest.ResponseRecorder
var mux *http.ServeMux

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Fatalf("Response code is %d", writer.Code)
	}

	actual, err := decodeTestData(writer)
	if err != nil {
		log.Fatal(err)
	}
	expected := 1

	if expected != actual.ID {
		t.Errorf("expected : %d, but was : %d", expected, actual.ID)
	}
}

func decodeTestData(writer *httptest.ResponseRecorder) (post Post, err error) {

	decoder := json.NewDecoder(writer.Body)

	for {
		err = decoder.Decode(&post)
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
