package main

import (
"net/http"
"net/http/httptest"
"testing"
	"strings"
	"net/url"
)

func Test_HelloWorld(t *testing.T) {
	//returns a new http request that can be used for testing
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	//ResponseRecorder is an implementation of http.ResponseWriter that records its mutations for later inspection in tests.
	//NewRecorder returns an initialized ResponseRecorder.
	res := httptest.NewRecorder()

	//use the request and the response and pass it to the handler
	HelloWorld(res, req)

	expected := "Hello World"
	actual := res.Body.String()
	if expected != actual {
		t.Fatalf("Expected %s gog %s", expected, actual)
	}
}

func Test_GenerateMarkdown(t *testing.T) {
	/* This type of body creation is need while doing a normal post body
	mcPostBody := map[string]string{
		"body": "Is this a test post for MutliQuestion?",
	}
	body, _ := json.Marshal(mcPostBody)*/

	//The type of form is url.values. It stores data in key value pairs
	form := url.Values{}
	form.Add("body","This is a test")
	req, err := http.NewRequest("POST", "www.example.com", strings.NewReader(form.Encode()))
	//req.ParseForm(form)
	req.PostForm = form
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	GenerateMarkdown(res,req)
	expected := "<p>This is a test</p>\n"
	actual := res.Body.String()

	if expected != actual {
		t.Fatalf("Expected : %s,%d ; got : %s, %d", expected,len(expected), actual, len(actual))
	}
}