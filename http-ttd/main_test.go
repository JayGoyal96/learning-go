package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBook(t *testing.T) {
	testcases := []struct {
		desc   string
		req    string
		expRes []Book
	}{
		{"get books", "/book", []Book{
			{"Jay", nil, "Jay", 2000},
			{"Goyal", nil, "Goyal", 1999},
		}},
		{},
	}
	for j, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "localhost:8000"+tc.req, nil)
		getBook(w, req)
		res, _ := io.ReadAll(w.Result().Body)
		resBooks := []Book{}
		json.Unmarshal(res, &resBooks)
		if len(resBooks) != len(tc.expRes) {
			t.Errorf("%v test failed %v", j, tc.desc)
		}
		for i := 0; i < len(resBooks); i++ {
			if resBooks[i] != tc.expRes[i] {
				t.Errorf("%v test failed %v", j, tc.desc)
			}
		}
	}
}

func TestGetBookById(t *testing.T) {
	testcases := []struct {
		desc   string
		req    string
		expRes Book
	}{
		{"get book", "/book/1", Book{
			Title:         "Jay",
			Author:        &Author{"jay", "Goyal", "11/03/2002", ""},
			Publication:   "",
			PublishedDate: 0,
		}},
		{},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "localhost:8000"+tc.req, nil)
		getBook(w, req)
		res, _ := io.ReadAll(w.Result().Body)
		resBook := Book{}
		json.Unmarshal(res, &resBook)
		if resBook != tc.expRes {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc    string
		req     string
		reqBody Book
		expRes  Book
	}{
		{"post book", "/book", Book{}, Book{}},
		{},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tc.reqBody)
		req := httptest.NewRequest(http.MethodPost, "localhost:8000"+tc.req, bytes.NewReader(body))
		postBook(w, req)
		res, _ := io.ReadAll(w.Result().Body)
		resBook := Book{}
		json.Unmarshal(res, &resBook)
		if resBook != tc.expRes {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc    string
		req     string
		reqBody Author
		expRes  Author
	}{
		{"post author", "/author", Author{}, Author{}},
		{},
	}
	for i, tc := range testcases {
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tc.reqBody)
		req := httptest.NewRequest(http.MethodPost, "localhost:8000"+tc.req, bytes.NewReader(body))
		postAuthor(w, req)
		res, _ := io.ReadAll(w.Result().Body)
		resAuthor := Author{}
		json.Unmarshal(res, &resAuthor)
		if resAuthor != tc.expRes {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}
