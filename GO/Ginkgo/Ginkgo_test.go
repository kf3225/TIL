package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "_/Users/k/TIL/GO/Ginkgo"
)

var _ = Describe("Ginkgo", func() {

	var mux *http.ServeMux
	var post *FakePost
	var writer *httptest.ResponseRecorder

	BeforeEach(func() {
		post = &FakePost{}
		mux = http.NewServeMux()
		writer = httptest.NewRecorder()

		mux.HandleFunc("/post/", HandleRequest(post))
	})

	Context("Get a post using an id", func() {
		It("should get a post", func() {
			request, _ := http.NewRequest(http.MethodGet, "/post/1", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(http.StatusOK))

			var post Post
			json.Unmarshal(writer.Body.Bytes(), &post)

			Expect(post.ID).To(Equal(1))
		})
	})

	Context("Get an error if id is not an Integer", func() {
		It("should get a HTTP 500 response", func() {
			request, _ := http.NewRequest(http.MethodGet, "/post/HelloWorld", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(500))
		})
	})
})
