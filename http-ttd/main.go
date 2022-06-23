package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Book struct {
	Title         string  `json:"title"`
	Author        *Author `json:"-"`
	Publication   string  `json:"publication"`
	PublishedDate int     `json:"published_date"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	PenName   string `json:"pen_name"`
}

func getBook(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode([]Book{{"Jay", nil, "Jay", 2000}, {"Goyal", nil, "Goyal", 1999}})
}

func getBookById(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(Book{
		Title: "Jay",
		Author: &Author{
			FirstName: "jay",
			LastName:  "Goyal",
			Dob:       "11/03/2002",
			PenName:   "",
		},
		Publication:   "",
		PublishedDate: 0,
	})
}

func postBook(response http.ResponseWriter, request *http.Request) {

}

func postAuthor(response http.ResponseWriter, request *http.Request) {

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/book", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			endpoints := strings.Split(request.URL.Path, "/")
			if len(endpoints) > 2 {
				fmt.Println("hello")
				getBookById(writer, request)
			} else {
				getBook(writer, request)
			}
		case http.MethodPost:
			endpoints := strings.Split(request.URL.Path, "/")
			if endpoints[1] == "book" {
				postBook(writer, request)
			} else {
				writer.WriteHeader(http.StatusBadRequest)
			}
		}
	})

	mux.HandleFunc("/author", postAuthor)

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	fmt.Println("Server started at 8000")
	server.ListenAndServe()
}
